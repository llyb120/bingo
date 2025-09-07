package ginx

import (
	"io"
	"net/http"
	"reflect"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/web"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var ParseJsonBodyNode = func(ctx core.Context, exist error) error {
	if exist != nil {
		return exist
	}
	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		return nil
	}
	// 只处理application/json
	if ginCtx.Request.Header.Get("Content-Type") != "application/json" {
		return nil
	}
	// 解析body
	body, err := io.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return err
	}
	// 保存原始值
	// ctx.Set("$bingo-body-raw", body)

	// 解析body到 T 或 *T
	_inType, ok := ctx.Get("$bingo-request-type")
	if !ok {
		return nil
	}
	inType, ok := _inType.(reflect.Type)
	if !ok {
		return nil
	}

	var parsed any
	if inType.Kind() == reflect.Ptr {
		// T 是指针类型：构造 *Elem，解码后得到 *Elem（即 T）
		v := reflect.New(inType.Elem())
		if err := json.Unmarshal(body, v.Interface()); err != nil {
			return err
		}
		parsed = v.Interface()
	} else {
		// T 是值类型：构造 *T 解码，再取值 T
		v := reflect.New(inType)
		if err := json.Unmarshal(body, v.Interface()); err != nil {
			return err
		}
		parsed = v.Elem().Interface()
	}
	// 保存解析后的值（类型与 T 对齐）
	ctx.Set("$bingo-body-parsed", parsed)

	return nil
}

var EvaluteServiceNode web.NodeHandler = nil

type Result struct {
	Data any    `json:"data"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var JsonResultNode = func(ctx core.Context, exist error) error {
	if exist != nil {
		return exist
	}
	c, ok := ctx.(*gin.Context)
	if !ok {
		return nil
	}

	// 获取解析后的值
	r, ok := ctx.Get("$bingo-service-result")
	if !ok {
		return nil
	}
	result := Result{
		Data: r,
		Code: 0,
		Msg:  "ok",
	}
	ctx.Set("$bingo-response", result)
	bs, err := json.Marshal(result)
	if err != nil {
		return err
	}
	c.Data(http.StatusOK, "application/json", bs)
	return nil
}

var ErrorResultNode = func(ctx core.Context, exist error) error {
	if exist == nil {
		return nil
	}
	c, ok := ctx.(*gin.Context)
	if !ok {
		return nil
	}
	r, _ := ctx.Get("$bingo-service-result")
	result := Result{
		Code: 1,
		Msg:  exist.Error(),
		Data: r,
	}
	ctx.Set("$bingo-response", result)
	bs, err := json.Marshal(result)
	if err != nil {
		return err
	}
	c.Data(http.StatusOK, "application/json", bs)
	return nil
}

type Validatable interface {
	Validate() error
}

var ValidateNode = func(ctx core.Context, exist error) error {
	if exist != nil {
		return exist
	}
	req, ok := ctx.Get("$bingo-body-parsed")
	if !ok {
		return nil
	}
	if v, ok := req.(Validatable); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type looseStringDecoder struct{}

func (d *looseStringDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.StringValue:
		*(*string)(ptr) = iter.ReadString()
	case jsoniter.NumberValue, jsoniter.BoolValue:
		*(*string)(ptr) = iter.ReadAny().ToString()
	case jsoniter.NilValue:
		iter.ReadNil()
		*(*string)(ptr) = ""
	default:
		*(*string)(ptr) = iter.ReadAny().ToString()
	}
}

type looseExt struct{ jsoniter.DummyExtension }

func (e *looseExt) CreateDecoder(t reflect2.Type) jsoniter.ValDecoder {
	if t.Kind() == reflect.String {
		return &looseStringDecoder{}
	}
	return nil
}

func init() {
	json.RegisterExtension(&looseExt{}) // 仅注册到上面的私有 json 实例
}
