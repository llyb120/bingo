package web

import (
	"github.com/llyb120/bingo/core"
)

type attachHandler[T any, R any] func(ctx core.Context, req T) (R, error)

type NodeHandler func(c core.Context, err error) error

type RequestHandler func(c core.Context)

//type Ctrl struct {
//	*CtrlContext
//	C    core.Context
//	Error error
//}
//
//type CtrlContext struct {
//	InType  reflect.Type
//	OutType reflect.Type
//}

func Attach[T any, R any](fn attachHandler[T, R]) func(c core.Context) {
	//controlContext := &CtrlContext{
	//	InType:  reflect.TypeOf((*T)(nil)).Elem(),
	//	OutType: reflect.TypeOf((*R)(nil)).Elem(),
	//}
	return func(c core.Context) {
		// c.Set("$bingo-ctrl-context", controlContext)
		//ctrl := &Ctrl{
		//	CtrlContext: controlContext,
		//	C:           c,
		//}
		// 运行handlers
		_handlers, ok := c.Get("$bingo-handlers")
		if !ok {
			return
		}
		handlers, ok := _handlers.([]NodeHandler)
		if !ok {
			return
		}
		var finalErr error
		for _, handler := range handlers {
			if handler == nil {
				if finalErr != nil {
					continue
				}
				_arg, _ := c.Get("$bingo-body-parsed")
				arg, ok := _arg.(T)
				if !ok {
					arg = *new(T)
				}
				r, err := fn(c, arg)
				if err != nil {
					finalErr = err
				}
				c.Set("$bingo-service-result", r)
				continue
			}

			//fn := handler(c)
			//if fn == nil {
			//	continue
			//}
			err := handler(c, finalErr) //fn()
			if err != nil {
				finalErr = err
				continue
				//c.JSON(http.StatusInternalServerError, gin.H{
				//	"error": ctrl.Error.Error(),
				//})
				//return
			}
		}
	}
}
