package core

import (
	"reflect"
	"sync"

	"github.com/petermattis/goid"
)

type Instance struct {
	Target any
	Name   string
}

type State struct {
	instanceMap map[any]*Instance
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

func newState() *State {
	return &State{
		instanceMap: make(map[any]*Instance),
		contextMap:  make(map[string]interface{}),
		bootPhase:   true,
		waitingFor:  make(map[int64]any),
		waitChans:   make(map[int64]chan struct{}),

		eventMap: make(map[string][]EventHandler),
	}
}

func (s *State) SetState(key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.contextMap[key] = value
}

func (s *State) GetState(key string) (any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.contextMap[key]
	return value, ok
}

// 开始boot阶段，记录总协程数
func (s *State) startBoot(totalGoroutines int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bootPhase = true
	s.totalGoroutines = totalGoroutines
}

// 结束boot阶段
func (s *State) endBoot() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.bootPhase = false
}

// Wire 从状态管理器获取实例并赋值给target
func (s *State) Use(target any, name ...string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		panic("target must be a pointer")
	}

	elemType := targetValue.Elem().Type()

	if len(name) == 0 {
		// 取符合条件的第一个
		for _, v := range s.instanceMap {
			if v.Target == nil {
				continue
			}

			targetVal := reflect.ValueOf(v.Target)
			// 1. 检查是否是相同类型
			// 2. 检查是否是指向相同类型的指针
			// 3. 检查是否实现了接口（如果目标类型是接口）
			if reflect.TypeOf(v.Target) == elemType ||
				(targetVal.Kind() == reflect.Ptr &&
					targetVal.Elem().Type() == elemType) ||
				(elemType.Kind() == reflect.Interface &&
					targetVal.Type().Implements(elemType)) {

				reflect.ValueOf(target).Elem().Set(reflect.ValueOf(v.Target))
				return
			}
		}
		// 取第一个
		panic("instance not found")
	} else {
		mp := s.instanceMap[name[0]]
		if mp == nil {
			panic("instance not found")
		}
		reflect.ValueOf(target).Elem().Set(reflect.ValueOf(mp.Target))
		return
	}
}

// Require 从状态管理器获取实例，带等待机制
func (s *State) Require(target any, name ...string) any {
	if !s.bootPhase {
		panic("require can only be called during boot/init phase")
	}

	// 检查依赖是否存在或注入目标
	s.mu.Lock()

	// 要求 target 必须是指针，用于注入
	targetVal := reflect.ValueOf(target)
	if targetVal.Kind() != reflect.Ptr {
		s.mu.Unlock()
		panic("target must be a pointer")
	}
	elemType := targetVal.Elem().Type()

	var instance *Instance
	if len(name) > 0 {
		// 按名称获取
		if inst, ok := s.instanceMap[name[0]]; ok {
			instance = inst
		}
	} else {
		// 按类型获取，与 Use 保持一致
		for _, inst := range s.instanceMap {
			if inst.Target == nil {
				continue
			}
			tVal := reflect.ValueOf(inst.Target)
			if reflect.TypeOf(inst.Target) == elemType ||
				(tVal.Kind() == reflect.Ptr && tVal.Elem().Type() == elemType) ||
				(elemType.Kind() == reflect.Interface && tVal.Type().Implements(elemType)) {
				instance = inst
				break
			}
		}
	}

	// 如果已找到实例，则直接注入并返回
	if instance != nil {
		reflect.ValueOf(target).Elem().Set(reflect.ValueOf(instance.Target))
		s.mu.Unlock()
		return instance.Target
	}

	// 依赖不存在，进入等待
	gid := goid.Get()

	var waitingKey any
	var instanceIdentifier string
	if len(name) > 0 {
		waitingKey = name[0]
		instanceIdentifier = name[0]
	} else {
		waitingKey = elemType
		instanceIdentifier = elemType.String()
	}

	// 检查是否已经处于等待状态（防止重入）
	if _, ok := s.waitChans[gid]; ok {
		s.mu.Unlock()
		panic("goroutine is already waiting for a dependency")
	}

	s.waitingFor[gid] = waitingKey
	waitChan := make(chan struct{})
	s.waitChans[gid] = waitChan
	s.waitingCount++

	// 检查死锁
	if s.totalGoroutines > 0 && s.waitingCount >= s.totalGoroutines {
		s.mu.Unlock()
		panic("circular dependency detected: all goroutines are waiting")
	}

	s.mu.Unlock()

	// 等待被唤醒
	<-waitChan

	// 被唤醒后，再次尝试获取
	s.mu.Lock()
	defer s.mu.Unlock()
	// 再次尝试获取实例
	var foundInstance *Instance
	if len(name) > 0 {
		if inst, ok := s.instanceMap[name[0]]; ok {
			foundInstance = inst
		}
	} else {
		for _, inst := range s.instanceMap {
			if inst.Target == nil {
				continue
			}
			tVal := reflect.ValueOf(inst.Target)
			if reflect.TypeOf(inst.Target) == elemType ||
				(tVal.Kind() == reflect.Ptr && tVal.Elem().Type() == elemType) ||
				(elemType.Kind() == reflect.Interface && tVal.Type().Implements(elemType)) {
				foundInstance = inst
				break
			}
		}
	}

	if foundInstance != nil {
		reflect.ValueOf(target).Elem().Set(reflect.ValueOf(foundInstance.Target))
		return foundInstance.Target
	}

	// 如果被唤醒后仍然找不到，说明有问题
	panic("instance still not found after waiting: " + instanceIdentifier)
}

// RegisterOption 注册选项
type RegisterOption struct {
	Name string
}
