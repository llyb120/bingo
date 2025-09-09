package main

import (
	"database/sql"
	"fmt"
	"github.com/llyb120/bingo/sqly/tool"

	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/log"
	"github.com/llyb120/bingo/sqly"
	"github.com/llyb120/bingo/web"
	"github.com/llyb120/bingo/web/ginx"
)

var RouterStarter core.Starter = func() func() {
	var r = core.Require[ginx.GinServer]()
	// 设定工作流
	r.AddNode(ginx.ParseJsonBodyNode, ginx.ValidateNode, ginx.EvaluteServiceNode, ginx.JsonResultNode, ginx.ErrorResultNode)
	// g.Use(parseBodyNode)

	core.On("before:sqly.Select", func(call *core.FuncCall) {
		// 示例：可根据 key 做缓存命中
		fmt.Println("intercept sql.Select")
		//call.Skip = true
		// 如果命中，可设置 call.Result 并 call.Skip = true 短路原调用
		_ = call
	})

	core.On("after:sqly.Select", func(call *core.FuncCall) {
		// 示例：将返回结果写入缓存
		call.Result = []any{[]map[string]string{
			{"a": "1"},
			{"a": "2"},
		}, nil}
		_ = call
	})

	r.GET("/", web.Attach(func(c core.Context, req struct {
		*TestValidate
		TopSize string `json:"top_size"`
	}) (any, error) {
		//cacheable := cache.Func_2_2(
		//	sqly.TestSelect,
		//	func(context core.Context, s string) string {
		//		return s
		//	},
		//	func(context core.Context, s string) time.Duration {
		//		return 10 * time.Second
		//	},
		//)
		//res, _ := cacheable(c, "ok")

		// 示例：查询返回原始数据（保持字段名）
		var db = core.Use[sql.DB]()
		// test sql template
		var sql, params, err = tool.GetSql("test.a")
		if err != nil {
			return nil, err
		}
		rawData, err := sqly.Select[[]map[string]any](c, db, sql, params...)
		if err != nil {
			return nil, err
		}
		log.Info(c, "sql is %s", sql)
		log.Info(c, "raw data: %v", rawData)

		rawData, err = sqly.Select[[]map[string]any](c, db, "select * from current_trade")
		if err != nil {
			return nil, err
		}

		// 示例：查询返回结构体数组（如果有定义相应结构体）
		// type Trade struct {
		//     ID int `json:"id"`
		//     Symbol string `json:"symbol"`
		//     Price float64 `json:"price"`
		// }
		// trades, err := sqly.Select[[]Trade](c, db, "select id, symbol, price from current_trade")

		// 示例：执行插入操作
		rowsAffected, lastInsertId, err := sqly.Exec(c, db, "INSERT INTO test_log (message, created_at) VALUES (?, NOW())", "API调用测试")
		if err != nil {
			log.Error(c, "插入日志失败: %v", err)
		}

		return map[string]interface{}{
			"raw_data":       rawData,
			"rows_affected":  rowsAffected,
			"last_insert_id": lastInsertId,
			"message":        "查询和插入操作完成",
		}, nil
	}))

	return nil
}
