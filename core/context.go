package core

import "context"

type Context interface {
	context.Context
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}
