package main

import (
	"database/sql"
	"fmt"
	"net/http"
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
	mysql0 *sql.DB
	mysql1 *sql.DB
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

	core.Use(&mysql0, "m0")
	core.Use(&mysql1)

	var gin ginx.GinServer
	core.Use(&gin)
	defer gin.Start()

	// 初始化路由
	initRouter(gin.Engine)

	fmt.Println(mysql0)
	fmt.Println(mysql1)
}

func initRouter(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		sqly.Select(mysql0, "select 1")
		c.String(http.StatusOK, "Hello, World!")
	})
}
