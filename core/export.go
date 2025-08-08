package core

import (
	"reflect"
)

func ExportFunc[T any](state *State, target *T, source T) {
	targetVal := reflect.ValueOf(target)
	sourceVal := reflect.ValueOf(source)

	// 检查 source 是否是函数
	if sourceVal.Kind() != reflect.Func {
		panic("source must be a function")
	}

	// 创建代理函数
	proxyFunc := func(in []reflect.Value) []reflect.Value {
		// 在实际调用前可以添加一些逻辑，比如日志、参数校验等
		// 调用原始函数
		results := sourceVal.Call(in)
		return results
	}

	// 创建与 source 函数相同类型的函数
	fnType := reflect.TypeOf(source)
	fn := reflect.MakeFunc(fnType, proxyFunc).Interface()

	// 将生成的函数赋值给 target
	targetElem := targetVal.Elem()
	if !targetElem.CanSet() {
		panic("target is not settable")
	}
	targetElem.Set(reflect.ValueOf(fn))
}

// ExportInstance 导出实例到状态管理器
func ExportInstance(s *State, instance any, args ...RegisterOption) {
	s.mu.Lock()

	var instanceName string
	// 如果提供了名字
	if len(args) > 0 && args[0].Name != "" {
		instanceName = args[0].Name
		s.instanceMap[instanceName] = &Instance{
			Target: instance,
			Name:   instanceName,
		}
	} else {
		// 否则使用自身类型
		typeOf := reflect.TypeOf(instance)
		instanceName = typeOf.String()
		s.instanceMap[typeOf] = &Instance{
			Target: instance,
			Name:   instanceName,
		}
	}

	// 唤醒等待此实例的goroutine
	exportedType := reflect.TypeOf(instance)
	for goid, waitingKey := range s.waitingFor {
		shouldWakeUp := false
		switch key := waitingKey.(type) {
		case string: // waiting for a name
			if key == instanceName {
				shouldWakeUp = true
			}
		case reflect.Type: // waiting for a type
			if exportedType == key ||
				(exportedType.Kind() == reflect.Ptr && exportedType.Elem() == key) ||
				(key.Kind() == reflect.Interface && exportedType.Implements(key)) {
				shouldWakeUp = true
			}
		}

		if shouldWakeUp {
			if ch, ok := s.waitChans[goid]; ok {
				close(ch)
				delete(s.waitingFor, goid)
				delete(s.waitChans, goid)
				s.waitingCount--
			}
		}
	}
	s.mu.Unlock()
}
