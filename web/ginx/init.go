package ginx

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
)

// var globalGin *gin.Engine
// var globalState *core.State

type GinServer struct {
	*gin.Engine
	config ginServer
}

var GinStarter core.Starter = func() func() {
	// globalState = state
	// globalGin = r
	r := gin.New()
	var server = ginServer{
		Port: 8080,
	}
	var config config.Config
	core.Require(&config)
	config.LoadToStruct("server", &server)

	core.ExportInstance(&GinServer{Engine: r, config: server}, core.RegisterOption{Name: "Gin"})
	return nil
}

type ginServer struct {
	Port int `json:"port"`
}

func (g *GinServer) Start() {
	go func() {
		g.Run(fmt.Sprintf(":%d", g.config.Port))
	}()

	quit := make(chan os.Signal)
	// sigint 是有 CTRL+C 触发，进程可以捕获信号，并进行处理，
	// SIGTERM是有 kill 命令触发,进程可以捕获信号，并进行处理，
	// SIGKILL是有 kill -9 命令触发，进程无法捕获。
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 处理后续逻辑
	fmt.Println("开始关闭gin server...")
}
