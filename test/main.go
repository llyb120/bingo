package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/llyb120/bingo/cache"
	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/datasource/mysql"
	"github.com/llyb120/bingo/datasource/redis"
	"github.com/llyb120/bingo/sqly"
	"github.com/llyb120/bingo/web/ginx"
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
	cache.CacheStarter,
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

func initRouter(g *ginx.GinServer) {

	// 设定工作流
	g.AddNode(ginx.ReadJsonBodyNode, ginx.ParseJsonBodyNode, ginx.EvaluteServiceNode, ginx.JsonResultNode)
	// g.Use(parseBodyNode)

	g.Engine.POST("/", ginx.Attach(func(ctx *gin.Context, req struct {
		TopSize string `json:"top_size"`
	}) (any, error) {
		sqly.Select(mysql0, "select 1")
		return "ok" + req.TopSize, nil
	}))
}
