package ginx

import (
	"io"
	"net/http"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var ReadJsonBodyNode = func(ctx *Ctrl) func() {
	// 只处理application/json
	if ctx.C.Request.Header.Get("Content-Type") != "application/json" {
		return nil
	}
	return func() {
		// 解析body
		body, err := io.ReadAll(ctx.C.Request.Body)
		if err != nil {
			ctx.Error = err
			return
		}
		// 保存原始值
		ctx.C.Set("$bingo-ginx-body-raw", body)
	}
}

var ParseJsonBodyNode = func(ctx *Ctrl) func() {
	// 只处理application/json
	if ctx.C.Request.Header.Get("Content-Type") != "application/json" {
		return nil
	}
	return func() {
		// 获取原始body
		body, ok := ctx.C.Get("$bingo-ginx-body-raw")
		if !ok {
			return
		}
		// 解析body到 T 或 *T
		inType := ctx.InType
		var parsed any
		if inType.Kind() == reflect.Ptr {
			// T 是指针类型：构造 *Elem，解码后得到 *Elem（即 T）
			v := reflect.New(inType.Elem())
			if err := json.Unmarshal(body.([]byte), v.Interface()); err != nil {
				ctx.Error = err
				return
			}
			parsed = v.Interface()
		} else {
			// T 是值类型：构造 *T 解码，再取值 T
			v := reflect.New(inType)
			if err := json.Unmarshal(body.([]byte), v.Interface()); err != nil {
				ctx.Error = err
				return
			}
			parsed = v.Elem().Interface()
		}
		// 保存解析后的值（类型与 T 对齐）
		ctx.C.Set("$bingo-ginx-body-parsed", parsed)
	}
}

var EvaluteServiceNode NodeHandler = nil

var JsonResultNode = func(ctx *Ctrl) func() {
	type Result struct {
		Data any    `json:"data"`
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	return func() {
		// 获取解析后的值
		r, ok := ctx.C.Get("$bingo-ginx-service-result")
		if !ok {
			return
		}
		result := Result{
			Data: r,
			Code: 0,
			Msg:  "ok",
		}
		ctx.C.Set("$bingo-ginx-response", result)
		bs, err := json.Marshal(result)
		if err != nil {
			ctx.Error = err
			return
		}
		ctx.C.Data(http.StatusOK, "application/json", bs)
	}
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
