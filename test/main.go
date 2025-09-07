package main

import (
	"database/sql"
	"fmt"
	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/datasource/mysql"
	"github.com/llyb120/bingo/datasource/redis"
	"github.com/llyb120/bingo/sqly"
	"github.com/llyb120/bingo/web"
	"github.com/llyb120/bingo/web/ginx"
	"os"
)

var (
	mysql0 = core.Use[sql.DB]("m0")
	mysql1 = core.Use[sql.DB]()
)

var plugins = []core.Starter{
	config.ConfigStarter,
	mysql.MysqlStarter,
	ginx.GinStarter,
	redis.RedisStarter,
	//cache.CacheStarter,
}

func init() {
	os.Setenv("BINGO_CONFIG_PATH", "./config.properties")
}

func main() {
	core.Boot(plugins...)
	defer core.Shutdown()

	// core.Use(&mysql0, "m0")
	// core.Use(&mysql1)

	var gin = core.Use[ginx.GinServer]()
	defer gin.Start()

	// 初始化路由
	initRouter(gin)
}

type TestValidate struct {
	TopSize string `json:"top_size"`
}

func (t *TestValidate) Validate() error {
	if t.TopSize == "" {
		return fmt.Errorf("top_size is empty")
	}
	return nil
}

func initRouter(g *ginx.GinServer) {

	// 设定工作流
	g.AddNode(ginx.ParseJsonBodyNode, ginx.ValidateNode, ginx.EvaluteServiceNode, ginx.JsonResultNode, ginx.ErrorResultNode)
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

	g.GET("/", web.Attach(func(c core.Context, req struct {
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

		res, err := sqly.Select(c, nil, "select * from test")
		return res, err
	}))
}
