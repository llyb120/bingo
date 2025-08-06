package core

import "sync"

type EventHandler func(args ...any)

var eventMap = make(map[string][]EventHandler)
var eventMapMutex sync.Mutex

func On(eventName string, callback EventHandler) {
	eventMapMutex.Lock()
	defer eventMapMutex.Unlock()
	eventMap[eventName] = append(eventMap[eventName], callback)
}

func Publish(eventName string, args ...any) {
	eventMapMutex.Lock()
	defer eventMapMutex.Unlock()
	for _, callback := range eventMap[eventName] {
		callback(args...)
	}
}
