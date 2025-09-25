package core

import (
	"context"
	"github.com/llyb120/yoya2/y"
)

type Context interface {
	context.Context
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type SimpleContext struct {
	context.Context
	mp map[string]interface{}
}

func (s *SimpleContext) Get(key string) (interface{}, bool) {
	v, ok := s.mp[key]
	return v, ok
}

func (s *SimpleContext) Set(key string, value interface{}) {
	s.mp[key] = value
}

var ContextPool = y.NewPool(func() *SimpleContext {
	return &SimpleContext{
		mp: make(map[string]interface{}),
	}
}, nil)
