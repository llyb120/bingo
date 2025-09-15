package core

import (
	"reflect"
	"sync"

	"github.com/petermattis/goid"
)

var globalState = newState()

type instance struct {
	Target any
	Name   string
}

type state struct {
	instanceMap map[any]*instance
	mu          sync.Mutex
	contextMap  map[string]interface{}

	// 依赖管理相关
	bootPhase       bool
	totalGoroutines int
	waitingCount    int
	// 记录哪个goid正在等待哪个实例 (string for name, reflect.Type for type)
	waitingFor map[int64]any
	// 记录等待goid的通知channel
	waitChans map[int64]chan struct{}

	// 事件订阅
	eventMap      map[string][]EventHandler
	eventMapMutex sync.RWMutex
}

func newState() *state {
	return &state{
		instanceMap: make(map[any]*instance),
		contextMap:  make(map[string]interface{}),
		bootPhase:   true,
		waitingFor:  make(map[int64]any),
		waitChans:   make(map[int64]chan struct{}),

		eventMap: make(map[string][]EventHandler),
	}
}

type Useable[T any] func() T

// tryAssignFromAny 支持以下注入场景（尽量避免反射，仅在指针解引用时使用）：
// 1) 直接断言到 T 成功（包含 T 为接口，且源实现该接口的情况）
// 2) 源为指针，且指向的元素可断言为 T（用于导出为 *U，目标为 U 的场景）
func tryAssignFromAny[T any](dst *T, src any) bool {
	if v, ok := src.(T); ok {
		*dst = v
		return true
	}
	rv := reflect.ValueOf(src)
	if rv.IsValid() && rv.Kind() == reflect.Ptr {
		ev := rv.Elem()
		if ev.IsValid() {
			if v2, ok := ev.Interface().(T); ok {
				*dst = v2
				return true
			}
		}
	}
	return false
}

func (s *state) SetState(key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.contextMap[key] = value
}

func (s *state) GetState(key string) (any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.contextMap[key]
	return value, ok
}

// 开始boot阶段，记录总协程数
func (s *state) startBoot(totalGoroutines int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bootPhase = true
	s.totalGoroutines = totalGoroutines
}

// 结束boot阶段
func (s *state) endBoot() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bootPhase = false
}

// Use 返回一个函数，该函数被调用时会从状态管理器获取实例。
// 这样避免了在 Use 调用时就需要处理注入时机的问题。
func Use[T any](name ...string) Useable[T] {
	return func() T {
		var s = globalState
		s.mu.Lock()
		defer s.mu.Unlock()

		// helper: 尝试从任意值中抽取 T
		getValue := func(src any) (T, bool) {
			// 直接类型断言
			if v, ok := src.(T); ok {
				return v, true
			}
			// 如果 src 是指针，尝试解引用
			rv := reflect.ValueOf(src)
			if rv.IsValid() && rv.Kind() == reflect.Ptr {
				ev := rv.Elem()
				if ev.IsValid() {
					if v, ok := ev.Interface().(T); ok {
						return v, true
					}
				}
			}
			var zero T
			return zero, false
		}

		// 命名注入优先
		if len(name) > 0 {
			if mp := s.instanceMap[name[0]]; mp != nil {
				if v, ok := getValue(mp.Target); ok {
					return v
				}
			}
			// 未找到时返回零值
			var zero T
			return zero
		}

		// 按类型/接口匹配：取符合条件的第一个
		for _, v := range s.instanceMap {
			if v == nil || v.Target == nil {
				continue
			}
			if val, ok := getValue(v.Target); ok {
				return val
			}
		}

		// 未找到时返回零值
		var zero T
		return zero
	}
}

// Require 从状态管理器获取实例，带等待机制。返回值类型与泛型参数一致。
func Require[T any](name ...string) T {
	var s = globalState
	if !s.bootPhase {
		panic("require can only be called during boot/init phase")
	}

	s.mu.Lock()
	// helper: 尝试从任意值中抽取 T
	getValue := func(src any) (T, bool) {
		// 直接类型断言
		if v, ok := src.(T); ok {
			return v, true
		}
		// 如果 src 是指针，尝试解引用
		rv := reflect.ValueOf(src)
		if rv.IsValid() && rv.Kind() == reflect.Ptr {
			ev := rv.Elem()
			if ev.IsValid() {
				if v, ok := ev.Interface().(T); ok {
					return v, true
				}
			}
		}
		var zero T
		return zero, false
	}

	// 先尝试立即获取
	if len(name) > 0 {
		if inst, ok := s.instanceMap[name[0]]; ok && inst != nil {
			if v, ok2 := getValue(inst.Target); ok2 {
				s.mu.Unlock()
				return v
			}
		}
	} else {
		for _, inst := range s.instanceMap {
			if inst == nil || inst.Target == nil {
				continue
			}
			if v, ok2 := getValue(inst.Target); ok2 {
				s.mu.Unlock()
				return v
			}
		}
	}

	// 未获取到，进入等待
	gid := goid.Get()

	var waitingKey any
	var instanceIdentifier string
	if len(name) > 0 {
		waitingKey = name[0]
		instanceIdentifier = name[0]
	} else {
		// 记录等待类型用于 ExportInstance 唤醒
		var zero T
		t := reflect.TypeOf(zero)
		waitingKey = t
		instanceIdentifier = t.String()
	}

	if _, ok := s.waitChans[gid]; ok {
		s.mu.Unlock()
		panic("goroutine is already waiting for a dependency")
	}

	s.waitingFor[gid] = waitingKey
	waitChan := make(chan struct{})
	s.waitChans[gid] = waitChan
	s.waitingCount++

	if s.totalGoroutines > 0 && s.waitingCount >= s.totalGoroutines {
		s.mu.Unlock()
		panic("circular dependency detected: all goroutines are waiting")
	}

	s.mu.Unlock()

	// 等待唤醒
	<-waitChan

	// 被唤醒后再次尝试
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(name) > 0 {
		if inst, ok := s.instanceMap[name[0]]; ok && inst != nil {
			if v, ok2 := getValue(inst.Target); ok2 {
				return v
			}
		}
	} else {
		for _, inst := range s.instanceMap {
			if inst == nil || inst.Target == nil {
				continue
			}
			if v, ok2 := getValue(inst.Target); ok2 {
				return v
			}
		}
	}

	panic("instance still not found after waiting: " + instanceIdentifier)
}

// RegisterOption 注册选项
type RegisterOption struct {
	Name string
}
