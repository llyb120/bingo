package ginx

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type attachHandler[T any, R any] func(ctx *gin.Context, req T) (R, error)

type NodeHandler func(c *Ctrl) func()

type Ctrl struct {
	*CtrlContext
	C     *gin.Context
	Error error
}

type CtrlContext struct {
	InType  reflect.Type
	OutType reflect.Type
}

func Attach[T any, R any](fn attachHandler[T, R]) func(c *gin.Context) {
	controlContext := &CtrlContext{
		InType:  reflect.TypeOf((*T)(nil)).Elem(),
		OutType: reflect.TypeOf((*R)(nil)).Elem(),
	}
	return func(c *gin.Context) {
		// c.Set("$bingo-ctrl-context", controlContext)
		ctrl := &Ctrl{
			CtrlContext: controlContext,
			C:           c,
		}
		// 运行handlers
		_handlers, ok := c.Get("$bingo-ginx-handlers")
		if !ok {
			return
		}
		handlers, ok := _handlers.([]NodeHandler)
		if !ok {
			return
		}
		for _, handler := range handlers {
			if handler == nil {
				_arg, _ := c.Get("$bingo-ginx-body-parsed")
				arg, ok := _arg.(T)
				if !ok {
					arg = *new(T)
				}
				r, err := fn(c, arg)
				if err != nil {
					ctrl.Error = err
					return
				}
				ctrl.C.Set("$bingo-ginx-service-result", r)
				continue
			}

			fn := handler(ctrl)
			if fn == nil {
				continue
			}
			fn()
			if ctrl.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": ctrl.Error.Error(),
				})
				return
			}
		}
	}
}
