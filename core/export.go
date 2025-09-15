package core

import (
	"reflect"
)

// FuncCall 表示一次通过 ExportFuncNamed 导出的函数调用
// 订阅者可在 before/after 事件中读取或修改调用
type FuncCall struct {
	state  *state
	Name   string
	Args   []any
	Result []any
	// 若在 before 阶段设置为 true，将跳过原函数调用，直接使用 Result 作为返回
	Skip bool
}

// ExportFuncNamed 通过 name 导出函数，并在调用前后发布事件：before:<name> / after:<name>
// 监听者可通过 *FuncCall 控制跳过原始调用或替换返回值，实现与调用方完全解耦
func ExportFunc[T any](name string, source T) T {
	sourceVal := reflect.ValueOf(source)

	if sourceVal.Kind() != reflect.Func {
		panic("source must be a function")
	}

	fnType := sourceVal.Type()

	// 将 []any 转换为目标函数签名的返回值切片
	toTypedReturns := func(results []any) []reflect.Value {
		out := make([]reflect.Value, fnType.NumOut())
		for i := 0; i < fnType.NumOut(); i++ {
			expected := fnType.Out(i)
			if i < len(results) && results[i] != nil {
				rv := reflect.ValueOf(results[i])
				if !rv.IsValid() {
					out[i] = reflect.Zero(expected)
				} else if rv.Type().AssignableTo(expected) {
					out[i] = rv
				} else if rv.Type().ConvertibleTo(expected) {
					out[i] = rv.Convert(expected)
				} else if expected.Kind() == reflect.Interface && rv.Type().Implements(expected) {
					out[i] = rv
				} else {
					out[i] = reflect.Zero(expected)
				}
			} else {
				out[i] = reflect.Zero(expected)
			}
		}
		return out
	}

	proxyFunc := func(in []reflect.Value) []reflect.Value {
		// 构造事件上下文
		args := make([]any, len(in))
		for i := range in {
			args[i] = in[i].Interface()
		}
		call := &FuncCall{Name: name, state: globalState, Args: args}

		// before
		Publish("before:"+name, call)
		if call.Skip {
			return toTypedReturns(call.Result)
		}

		// 原始调用
		results := sourceVal.Call(in)

		// after（允许替换返回值）
		resAny := make([]any, len(results))
		for i := range results {
			resAny[i] = results[i].Interface()
		}
		call.Result = resAny
		Publish("after:"+name, call)

		// 如 after 修改了 Result，则以修改后的为准
		if call.Result != nil {
			return toTypedReturns(call.Result)
		}
		return results
	}

	fn := reflect.MakeFunc(fnType, proxyFunc).Interface()
	return fn.(T)
}

// ExportInstance 导出实例到状态管理器
func ExportInstance(ins any, args ...RegisterOption) {
	s := globalState
	s.mu.Lock()

	var instanceName string
	// 统一将实例以“指针”形式存储，确保后续 Use/Require 可返回原始指针
	stored := ins
	rv := reflect.ValueOf(ins)
	if rv.IsValid() && rv.Kind() != reflect.Ptr {
		ptr := reflect.New(rv.Type())
		ptr.Elem().Set(rv)
		stored = ptr.Interface()
	}

	// 如果提供了名字
	if len(args) > 0 && args[0].Name != "" {
		instanceName = args[0].Name
		s.instanceMap[instanceName] = &instance{
			Target: stored,
			Name:   instanceName,
		}
	} else {
		// 否则使用自身类型
		typeOf := reflect.TypeOf(stored)
		instanceName = typeOf.String()
		s.instanceMap[typeOf] = &instance{
			Target: stored,
			Name:   instanceName,
		}
	}

	// 唤醒等待此实例的goroutine（针对 Require）
	exportedType := reflect.TypeOf(ins)
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

// func Around[T any](state *State, sourceFn T, fn func(next T)) {
// 	state.aroundConfig.aroundMu.Lock()
// 	// 以原始函数的指针唯一标识（仅 ExportFunc 导出的函数可被注册）
// 	key := reflect.ValueOf(sourceFn).Pointer()
// 	state.aroundConfig.aroundMap[key] = append(state.aroundConfig.aroundMap[key], fn)
// 	state.aroundConfig.aroundMu.Unlock()
// }
