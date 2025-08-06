package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Config 配置结构体
type Config map[string]interface{}

var (
	arrayKeyRegexp = regexp.MustCompile(`^(\w+)\[(\d+)\]$`)
)

func (c Config) LoadProperties(path string) (Config, error) {
	return loadProperties(path)
}

func mustLoadProperties(path string) Config {
	cfg, err := loadProperties(path)
	if err != nil {
		// 配置文件读取失败时返回空配置，不再panic
		return make(Config)
	}
	return cfg
}

func (c Config) LoadToStruct(prefix string, cfg interface{}) {
	if err := c.loadToStruct(prefix, cfg); err != nil {
		panic(err)
	}
}

// LoadProperties 加载 properties 文件
// 支持两种格式：
// 1. 数组格式: datasource.mysql[0].host=127.0.0.1
// 2. Map格式: datasource.mysql.abc.host=127.0.0.1
func loadProperties(path string) (Config, error) {
	cfg := make(Config)

	// 获取绝对路径
	abspath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	// 打开文件
	file, err := os.Open(abspath)
	if err != nil {
		// 文件不存在时返回空配置，不报错
		return cfg, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// 跳过空行和注释
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "!") {
			continue
		}

		// 解析键值对
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format at line %d: %s", lineNum, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// 处理键中的数组索引（如 "datasource.mysql[0].host"）
		keys := strings.Split(key, ".")
		current := cfg

		for i := 0; i < len(keys)-1; i++ {
			k := keys[i]

			// 检查是否是数组键（如 "mysql[0]"）
			if matches := arrayKeyRegexp.FindStringSubmatch(k); len(matches) == 3 {
				arrayName := matches[1]
				arrayIndexStr := matches[2]
				arrayIndex, _ := strconv.Atoi(arrayIndexStr)

				// 如果当前节点不存在，初始化为数组
				if _, exists := current[arrayName]; !exists {
					current[arrayName] = make([]interface{}, arrayIndex+1)
				}

				array, ok := current[arrayName].([]interface{})
				if !ok {
					// 如果不是数组，则将数字索引作为map的key
					mapVal, isMap := current[arrayName].(Config)
					if !isMap {
						mapVal = make(Config)
						current[arrayName] = mapVal
					}

					// 将数字索引作为map的key
					if _, exists := mapVal[arrayIndexStr]; !exists {
						mapVal[arrayIndexStr] = make(Config)
					}

					next, ok := mapVal[arrayIndexStr].(Config)
					if !ok {
						// 如果已存在但不是Config类型，覆盖为Config
						mapVal[arrayIndexStr] = make(Config)
						next = mapVal[arrayIndexStr].(Config)
					}
					current = next
					continue
				}

				// 扩展数组大小（如果需要）
				if arrayIndex >= len(array) {
					newArray := make([]interface{}, arrayIndex+1)
					copy(newArray, array)
					array = newArray
					current[arrayName] = array
				}

				// 如果当前位置是nil，初始化为新的 Config
				if array[arrayIndex] == nil {
					array[arrayIndex] = make(Config)
				}

				// 检查类型是否正确
				next, ok := array[arrayIndex].(Config)
				if !ok {
					// 如果不是Config类型，尝试将其转换为Config
					array[arrayIndex] = make(Config)
					next = array[arrayIndex].(Config)
				}

				current = next
			} else {
				// 普通键处理
				if _, exists := current[k]; !exists {
					current[k] = make(Config)
				} else if _, isMap := current[k].(Config); !isMap {
					return nil, fmt.Errorf("key conflict at %s: not a section", strings.Join(keys[:i+1], "."))
				}
				current = current[k].(Config)
			}
		}

		// 设置值
		lastKey := keys[len(keys)-1]
		if matches := arrayKeyRegexp.FindStringSubmatch(lastKey); len(matches) == 3 {
			// 处理最后一个键是数组的情况
			arrayName := matches[1]
			arrayIndexStr := matches[2]
			arrayIndex, _ := strconv.Atoi(arrayIndexStr)

			if _, exists := current[arrayName]; !exists {
				current[arrayName] = make([]interface{}, arrayIndex+1)
			}

			array, ok := current[arrayName].([]interface{})
			if !ok {
				// 如果不是数组，则将数字索引作为map的key
				mapVal, isMap := current[arrayName].(Config)
				if !isMap {
					mapVal = make(Config)
					current[arrayName] = mapVal
				}
				mapVal[arrayIndexStr] = value
				continue
			}

			if arrayIndex >= len(array) {
				newArray := make([]interface{}, arrayIndex+1)
				copy(newArray, array)
				array = newArray
				current[arrayName] = array
			}

			array[arrayIndex] = value
		} else {
			// 普通键设置值
			current[lastKey] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// 应用环境变量覆盖
	// 暂时不处理环境变量，避免冲突
	return cfg, nil
}

// GetSection 获取配置的子节
func (c Config) GetSection(key string) Config {
	if key == "" {
		return c
	}

	keys := strings.Split(key, ".")
	current := c

	for _, k := range keys {
		val, exists := current[k]
		if !exists {
			return nil
		}

		if next, ok := val.(Config); ok {
			current = next
		} else {
			return nil
		}
	}

	return current
}

// GetArray 获取数组配置
// 例如：datasource.mysql[0].host=127.0.0.1
func (c Config) GetArray(key string) []interface{} {
	if key == "" {
		return nil
	}

	keys := strings.Split(key, ".")
	current := c

	// 遍历到倒数第二个键
	for i := 0; i < len(keys)-1; i++ {
		val, exists := current[keys[i]]
		if !exists {
			return nil
		}

		if next, ok := val.(Config); ok {
			current = next
		} else {
			return nil
		}
	}

	// 获取最后一个键对应的值
	lastKey := keys[len(keys)-1]
	val, exists := current[lastKey]
	if !exists {
		return nil
	}

	arr, ok := val.([]interface{})
	if !ok {
		return nil
	}

	return arr
}

// GetMap 获取Map配置
// 例如：datasource.mysql.abc.host=127.0.0.1
func (c Config) GetMap(key string) map[string]interface{} {
	if key == "" {
		return nil
	}

	keys := strings.Split(key, ".")
	current := c

	// 遍历到倒数第二个键
	for i := 0; i < len(keys)-1; i++ {
		val, exists := current[keys[i]]
		if !exists {
			return nil
		}

		if next, ok := val.(Config); ok {
			current = next
		} else {
			return nil
		}
	}

	// 获取最后一个键对应的值
	lastKey := keys[len(keys)-1]
	val, exists := current[lastKey]
	if !exists {
		return nil
	}

	result := make(map[string]interface{})
	if config, ok := val.(Config); ok {
		for k, v := range config {
			result[k] = v
		}
		return result
	}

	return nil
}

// LoadToStruct 加载配置到指定结构体
// prefix 参数用于指定配置的前缀，例如 "datasource.mysql" 会加载 datasource.mysql 开头的配置
// cfg 必须是一个结构体指针或结构体指针的切片
// 支持两种格式：
// 1. 数组格式: datasource.mysql[0].host=127.0.0.1
// 2. Map格式: datasource.mysql.abc.host=127.0.0.1
func (c Config) loadToStruct(prefix string, cfg interface{}) error {
	v := reflect.ValueOf(cfg)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("config must be a non-nil pointer to struct or slice of struct pointers")
	}

	v = v.Elem()

	// 处理切片类型
	if v.Kind() == reflect.Slice {
		// 获取切片元素类型
		elemType := v.Type().Elem()
		if elemType.Kind() != reflect.Ptr || elemType.Elem().Kind() != reflect.Struct {
			return fmt.Errorf("slice elements must be pointers to struct")
		}

		// 获取前缀对应的配置
		var configData interface{} = c
		if prefix != "" {
			keys := strings.Split(prefix, ".")
			for _, key := range keys {
				if m, ok := configData.(Config); ok {
					if val, exists := m[key]; exists {
						configData = val
					} else {
						// 配置读不出来返回空值，不报错
						return nil
					}
				} else {
					// 配置读不出来返回空值，不报错
					return nil
				}
			}
		}

		// 处理数组配置
		if arr, ok := configData.([]interface{}); ok {
			slice := reflect.MakeSlice(v.Type(), len(arr), len(arr))
			for i, item := range arr {
				if item == nil {
					continue
				}

				// 创建新的结构体实例
				elem := reflect.New(elemType.Elem())

				// 递归处理结构体字段
				if m, ok := item.(map[string]interface{}); ok {
					if err := setStructFields(Config(m), elem.Elem()); err != nil {
						return fmt.Errorf("failed to set struct fields at index %d: %v", i, err)
					}
				} else if cfg, ok := item.(Config); ok {
					if err := setStructFields(cfg, elem.Elem()); err != nil {
						return fmt.Errorf("failed to set struct fields at index %d: %v", i, err)
					}
				}

				slice.Index(i).Set(elem)
			}

			v.Set(slice)
			return nil
		}

		// 配置读不出来返回空值，不报错
		return nil
	}

	// 处理 map 类型
	if v.Kind() == reflect.Map {
		if v.Type().Key().Kind() != reflect.String {
			return fmt.Errorf("map keys must be strings")
		}

		// 获取 map 的值类型
		elemType := v.Type().Elem()
		var isPtr bool

		if elemType.Kind() == reflect.Struct {
			// 结构体类型
		} else if elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct {
			isPtr = true
		} else {
			return fmt.Errorf("map values must be structs or pointers to struct")
		}

		// 获取前缀对应的配置
		current := c

		if prefix != "" {
			keys := strings.Split(prefix, ".")
			for _, key := range keys {
				if val, exists := current[key]; exists {
					if cfg, ok := val.(Config); ok {
						current = cfg
					} else {
						// 配置读不出来返回空值，不报错
						return nil
					}
				} else {
					// 配置读不出来返回空值，不报错
					return nil
				}
			}
		}

		// 处理 map 配置
		result := reflect.MakeMap(v.Type())
		for key, value := range current {
			if value == nil {
				continue
			}

			var elem reflect.Value
			if isPtr {
				elem = reflect.New(elemType.Elem())
			} else {
				elem = reflect.New(elemType).Elem()
			}

			// 递归处理结构体字段
			if cfg, ok := value.(Config); ok {
				if err := setStructFields(cfg, elem); err != nil {
					return fmt.Errorf("failed to set struct fields for key %s: %v", key, err)
				}
			} else if m, ok := value.(map[string]interface{}); ok {
				if err := setStructFields(Config(m), elem); err != nil {
					return fmt.Errorf("failed to set struct fields for key %s: %v", key, err)
				}
			}

			if isPtr {
				result.SetMapIndex(reflect.ValueOf(key), elem)
			} else {
				result.SetMapIndex(reflect.ValueOf(key), elem)
			}
		}

		v.Set(result)
		return nil
	}

	// 处理结构体类型
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("config must be a pointer to struct, slice of struct pointers, or map of structs")
	}

	current := c

	// 获取前缀对应的配置
	if prefix != "" {
		keys := strings.Split(prefix, ".")
		for _, key := range keys {
			if val, exists := current[key]; exists {
				if cfg, ok := val.(Config); ok {
					current = cfg
				} else {
					// 配置读不出来返回空值，不报错
					return nil
				}
			} else {
				// 配置读不出来返回空值，不报错
				return nil
			}
		}
	}

	// 设置结构体字段
	return setStructFields(current, v)
}

// applyEnvOverrides 应用环境变量覆盖配置
func applyEnvOverrides(cfg Config, prefix string) (Config, error) {
	// 创建配置的副本
	result := make(Config)
	for k, v := range cfg {
		if subCfg, ok := v.(Config); ok {
			subResult, err := applyEnvOverrides(subCfg, prefix+k+"_")
			if err != nil {
				return nil, err
			}
			result[k] = subResult
		} else {
			result[k] = v
		}
	}

	// 应用环境变量覆盖
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) != 2 {
			continue
		}

		envKey := parts[0]
		envValue := parts[1]

		// 如果指定了前缀，只处理以该前缀开头的环境变量
		if prefix != "" && !strings.HasPrefix(envKey, prefix) {
			continue
		}

		// 移除前缀并转换为小写，将下划线替换为点
		key := strings.TrimPrefix(envKey, prefix)
		key = strings.ToLower(key)
		key = strings.ReplaceAll(key, "_", ".")

		// 设置配置值
		keys := strings.Split(key, ".")
		current := result

		for i, k := range keys[:len(keys)-1] {
			if _, exists := current[k]; !exists {
				current[k] = make(Config)
			} else if _, isMap := current[k].(Config); !isMap {
				return nil, fmt.Errorf("key conflict at %s: not a section", strings.Join(keys[:i+1], "."))
			}
			current = current[k].(Config)
		}

		current[keys[len(keys)-1]] = envValue
	}

	return result, nil
}

// setStructFields 递归设置结构体字段
func setStructFields(cfg Config, v reflect.Value) error {
	// 处理指针类型
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct or pointer to struct, got %s", v.Kind())
	}

	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		// 获取字段的 prop 标签
		tag := field.Tag.Get("json")
		if tag == "" {
			tag = strings.ToLower(field.Name)
		}

		// 如果是嵌套结构体，递归处理
		if field.Type.Kind() == reflect.Struct {
			subCfg, exists := cfg[tag]
			if !exists {
				continue
			}

			subMap, ok := subCfg.(Config)
			if !ok {
				return fmt.Errorf("field %s is not a config section", field.Name)
			}

			if err := setStructFields(subMap, fieldVal); err != nil {
				return err
			}
			continue
		}

		// 处理指针类型的嵌套结构体
		if field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct {
			subCfg, exists := cfg[tag]
			if !exists {
				continue
			}

			subMap, ok := subCfg.(Config)
			if !ok {
				return fmt.Errorf("field %s is not a config section", field.Name)
			}

			elem := reflect.New(field.Type.Elem())
			if err := setStructFields(subMap, elem.Elem()); err != nil {
				return err
			}
			fieldVal.Set(elem)
			continue
		}

		// 获取配置值
		val, exists := cfg[tag]
		if !exists {
			continue
		}

		switch field.Type.Kind() {
		case reflect.String:
			fieldVal.SetString(fmt.Sprint(val))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			switch v := val.(type) {
			case int:
				fieldVal.SetInt(int64(v))
			case float64:
				fieldVal.SetInt(int64(v))
			case string:
				if i, err := strconv.ParseInt(v, 10, 64); err == nil {
					fieldVal.SetInt(i)
				}
			}
		case reflect.Bool:
			switch v := val.(type) {
			case bool:
				fieldVal.SetBool(v)
			case string:
				fieldVal.SetBool(strings.ToLower(v) == "true" || v == "1")
			}
		case reflect.Float32, reflect.Float64:
			switch v := val.(type) {
			case float64:
				fieldVal.SetFloat(v)
			case string:
				if f, err := strconv.ParseFloat(v, 64); err == nil {
					fieldVal.SetFloat(f)
				}
			}
		}
	}

	return nil
}

// GetString 获取字符串配置值
func (c Config) GetString(key string) string {
	val := c.getValue(key)
	if val == nil {
		return ""
	}
	return fmt.Sprint(val)
}

// GetInt 获取整数配置值
func (c Config) GetInt(key string) int {
	val := c.getValue(key)
	if val == nil {
		return 0
	}

	switch v := val.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case string:
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return 0
}

// GetBool 获取布尔配置值
func (c Config) GetBool(key string) bool {
	val := c.getValue(key)
	if val == nil {
		return false
	}

	switch v := val.(type) {
	case bool:
		return v
	case string:
		return strings.ToLower(v) == "true" || v == "1"
	}
	return false
}

// GetFloat 获取浮点数配置值
func (c Config) GetFloat(key string) float64 {
	val := c.getValue(key)
	if val == nil {
		return 0.0
	}

	switch v := val.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return 0.0
}

// getValue 获取配置值的内部方法
func (c Config) getValue(key string) interface{} {
	if key == "" {
		return nil
	}

	keys := strings.Split(key, ".")
	current := c

	for i, k := range keys {
		val, exists := current[k]
		if !exists {
			return nil
		}

		// 如果是最后一个键，直接返回值
		if i == len(keys)-1 {
			return val
		}

		// 继续深入下一层
		if next, ok := val.(Config); ok {
			current = next
		} else {
			return nil
		}
	}

	return nil
}
