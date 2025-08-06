package core

import (
	"reflect"
)

func (s *State) ExportFunc(target any, source any) {
	targetVal := reflect.ValueOf(target)
	sourceVal := reflect.ValueOf(source)

	// 检查 target 是否是指针的指针
	if targetVal.Kind() != reflect.Ptr || targetVal.Elem().Kind() != reflect.Func {
		panic("target must be a pointer to a function")
	}

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
