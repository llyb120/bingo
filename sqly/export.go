package sqly

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"

	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/log"
	"github.com/llyb120/yoya2/y"
)

//var (
//	Select = core.ExportFunc("sqly.Select", _Select)
//)

// Select 执行查询并转换成指定类型T的结果返回
func Select[T any](ctx core.Context, db DBLike, sql string, args ...any) (T, error) {
	var zero T

	log.Info(ctx, "执行SQL查询: %s", sql)

	if db == nil {
		return zero, fmt.Errorf("数据库连接不能为空")
	}

	// 执行查询获取map格式数据（保持字段名）
	data, err := queryToMaps(ctx, db, sql, args...)
	if err != nil {
		return zero, err
	}

	// 使用y.Cast转换成目标类型
	var result T
	err = y.Cast(&result, data)
	if err != nil {
		log.Error(ctx, "类型转换失败: %v", err)
		return zero, fmt.Errorf("类型转换失败: %w", err)
	}

	return result, nil
}

// Exec 执行SQL语句（INSERT、UPDATE、DELETE等）并返回影响的行数、自增ID、错误
func Exec(ctx core.Context, db DBLike, sql string, args ...any) (int64, int64, error) {
	log.Info(ctx, "执行SQL语句: %s", sql)

	if db == nil {
		return 0, 0, fmt.Errorf("数据库连接不能为空")
	}

	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Error(ctx, "SQL执行失败: %v", err)
		return 0, 0, fmt.Errorf("SQL执行失败: %w", err)
	}

	// 获取受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error(ctx, "获取影响行数失败: %v", err)
		return 0, 0, fmt.Errorf("获取影响行数失败: %w", err)
	}

	// 获取自增ID（如果有）
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		// 某些数据库驱动或操作不支持LastInsertId，这是正常的
		log.Debug(ctx, "获取自增ID失败（可能不支持或不适用）: %v", err)
		lastInsertId = 0 // 设置为0表示没有自增ID或不支持
	}

	log.Info(ctx, "SQL执行成功，影响行数: %d，自增ID: %d", rowsAffected, lastInsertId)
	return rowsAffected, lastInsertId, nil
}

// queryToMaps 执行查询并返回[]map[string]any格式（保持字段名）
func queryToMaps(ctx core.Context, db DBLike, sql string, args ...any) ([]map[string]any, error) {
	// 执行查询
	rows, err := db.Query(sql, args...)
	if err != nil {
		log.Error(ctx, "查询执行失败: %v", err)
		return nil, fmt.Errorf("查询执行失败: %w", err)
	}
	defer rows.Close()

	// 获取列信息
	columns, err := rows.Columns()
	if err != nil {
		log.Error(ctx, "获取列信息失败: %v", err)
		return nil, fmt.Errorf("获取列信息失败: %w", err)
	}

	// 获取列类型信息
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Error(ctx, "获取列类型失败: %v", err)
		return nil, fmt.Errorf("获取列类型失败: %w", err)
	}

	var result []map[string]any

	// 遍历所有行
	for rows.Next() {
		// 创建扫描目标
		scanDest := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanDest[i] = &values[i]
		}

		// 扫描行
		if err := rows.Scan(scanDest...); err != nil {
			log.Error(ctx, "扫描行失败: %v", err)
			return nil, fmt.Errorf("扫描行失败: %w", err)
		}

		// 转换成map[string]any格式（保持字段名）
		rowMap := make(map[string]any)
		for i, v := range values {
			rowMap[columns[i]] = convertToGoType(v, columnTypes[i])
		}

		result = append(result, rowMap)
	}

	// 检查是否有错误
	if err = rows.Err(); err != nil {
		log.Error(ctx, "行扫描错误: %v", err)
		return nil, fmt.Errorf("行扫描错误: %w", err)
	}

	return result, nil
}

// convertToGoType 将数据库驱动返回的值转换为Go基本类型
func convertToGoType(value interface{}, columnType *sql.ColumnType) interface{} {
	if value == nil {
		return nil
	}

	// 获取数据库类型名称
	dbTypeName := columnType.DatabaseTypeName()

	// 处理不同的数据库类型（先处理具体类型，再处理接口类型）
	switch v := value.(type) {
	case []byte:
		// 字节数组通常需要转换为字符串或数字
		return convertBytesToGoType(v, dbTypeName)
	case sql.NullString:
		if v.Valid {
			return v.String
		}
		return nil
	case sql.NullInt64:
		if v.Valid {
			return v.Int64
		}
		return nil
	case sql.NullInt32:
		if v.Valid {
			return v.Int32
		}
		return nil
	case sql.NullFloat64:
		if v.Valid {
			return v.Float64
		}
		return nil
	case sql.NullBool:
		if v.Valid {
			return v.Bool
		}
		return nil
	case sql.NullTime:
		if v.Valid {
			return v.Time
		}
		return nil
	case time.Time:
		return v
	case int, int8, int16, int32, int64:
		return v
	case uint, uint8, uint16, uint32, uint64:
		return v
	case float32, float64:
		return v
	case bool:
		return v
	case string:
		return v
	case driver.Valuer:
		// 如果是driver.Valuer，先获取其值
		if val, err := v.Value(); err == nil {
			return convertToGoType(val, columnType)
		}
		return v
	default:
		// 对于其他类型，尝试反射获取底层值
		return convertWithReflection(value, dbTypeName)
	}
}

// convertBytesToGoType 根据数据库类型将字节数组转换为合适的Go类型
func convertBytesToGoType(data []byte, dbType string) interface{} {
	if len(data) == 0 {
		return nil
	}

	str := string(data)

	// 根据数据库类型判断目标类型
	switch dbType {
	case "TINYINT", "SMALLINT", "MEDIUMINT", "INT", "BIGINT":
		// 整数类型
		if num, err := parseNumber(str); err == nil {
			return num
		}
		return str
	case "DECIMAL", "NUMERIC", "FLOAT", "DOUBLE":
		// 浮点数类型
		if num, err := parseFloat(str); err == nil {
			return num
		}
		return str
	case "BIT", "BOOLEAN":
		// 布尔类型
		if b, err := parseBool(str); err == nil {
			return b
		}
		return str
	case "DATE", "TIME", "DATETIME", "TIMESTAMP":
		// 时间类型
		if t, err := parseTime(str); err == nil {
			return t
		}
		return str
	default:
		// 默认返回字符串
		return str
	}
}

// convertWithReflection 使用反射尝试转换未知类型
func convertWithReflection(value interface{}, dbType string) interface{} {
	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return nil
	}

	// 如果是指针，获取指向的值
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}

	// 返回底层接口值
	if rv.CanInterface() {
		return rv.Interface()
	}

	return value
}

// 辅助函数：解析数字
func parseNumber(s string) (interface{}, error) {
	// 尝试解析为int64
	if num, err := fmt.Sscanf(s, "%d", new(int64)); err == nil && num == 1 {
		var result int64
		fmt.Sscanf(s, "%d", &result)
		return result, nil
	}
	return nil, fmt.Errorf("无法解析为数字: %s", s)
}

// 辅助函数：解析浮点数
func parseFloat(s string) (interface{}, error) {
	if num, err := fmt.Sscanf(s, "%f", new(float64)); err == nil && num == 1 {
		var result float64
		fmt.Sscanf(s, "%f", &result)
		return result, nil
	}
	return nil, fmt.Errorf("无法解析为浮点数: %s", s)
}

// 辅助函数：解析布尔值
func parseBool(s string) (interface{}, error) {
	switch s {
	case "1", "true", "TRUE", "True", "yes", "YES", "Yes":
		return true, nil
	case "0", "false", "FALSE", "False", "no", "NO", "No":
		return false, nil
	}
	return nil, fmt.Errorf("无法解析为布尔值: %s", s)
}

// 辅助函数：解析时间
func parseTime(s string) (interface{}, error) {
	// 尝试多种时间格式
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02",
		"15:04:05",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, s); err == nil {
			return t, nil
		}
	}

	return nil, fmt.Errorf("无法解析为时间: %s", s)
}

func TestSelect(ctx core.Context, ok string) (string, error) {
	return ok + "fu", nil
}
