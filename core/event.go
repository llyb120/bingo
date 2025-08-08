package core

type EventHandler func(args ...any)

func (s *State) On(eventName string, callback EventHandler) {
	s.eventMapMutex.Lock()
	defer s.eventMapMutex.Unlock()
	s.eventMap[eventName] = append(s.eventMap[eventName], callback)
}

func (s *State) Publish(eventName string, args ...any) {
	// 读取时加读锁，复制切片后释放锁，降低锁粒度
	s.eventMapMutex.RLock()
	handlers := append([]EventHandler(nil), s.eventMap[eventName]...)
	s.eventMapMutex.RUnlock()
	for _, callback := range handlers {
		callback(args...)
	}
}
