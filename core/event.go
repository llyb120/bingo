package core

type EventHandler func(args *FuncCall)

func On(eventName string, callback EventHandler) {
	var s = globalState
	s.eventMapMutex.Lock()
	defer s.eventMapMutex.Unlock()
	s.eventMap[eventName] = append(s.eventMap[eventName], callback)
}

func Publish(eventName string, args *FuncCall) {
	var s = globalState
	// 读取时加读锁，复制切片后释放锁，降低锁粒度
	s.eventMapMutex.RLock()
	handlers := append([]EventHandler(nil), s.eventMap[eventName]...)
	s.eventMapMutex.RUnlock()
	for _, callback := range handlers {
		callback(args)
	}
}
