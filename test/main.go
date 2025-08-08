package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/datasource/mysql"
	"github.com/llyb120/bingo/datasource/redis"
	"github.com/llyb120/bingo/web/ginx"
)

var (
	mysql0 *sql.DB
	mysql1 *sql.DB
)

func main() {
	os.Setenv("BINGO_CONFIG_PATH", "./config.properties")
	state := core.Boot(config.ConfigStarter, mysql.MysqlStarter, ginx.GinStarter, redis.RedisStarter)
	defer state.Shutdown()
	state.Use(&mysql0, "m0")
	state.Use(&mysql1)

	var gin *gin.Engine
	state.Use(&gin)

	defer ginx.Start(state)

	// 初始化路由
	initRouter(gin)

	fmt.Println(mysql0)
	fmt.Println(mysql1)
}

func initRouter(g *gin.Engine) {
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
}
