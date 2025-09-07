package ginx

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/llyb120/bingo/web"

	"github.com/gin-gonic/gin"
)

var (
// _ web.WebServer = (*GinServer)(nil)
)

type GinServer struct {
	Engine   *gin.Engine
	config   ginServer
	handlers []web.NodeHandler
}

// 实现 web.Server
// func (g *GinServer) Get(path string, handler func(ctx *web.Context)) {
// 	g.Engine.GET(path, func(c *gin.Context) {
// 		handler(web.Context{
// 			mp: c.GetMap(),
// 		})
// 	})
// }

// func (g *GinServer) Post(path string, handler func(ctx *web.Context)) {
// 	g.Engine.POST(path, func(c *gin.Context) {
// 		// 申请 context
// 		ctx, recycle := web.ContextPool.Get()
// 		defer recycle()
// 		// 包装
// 		web.Wrap(ctx, c)
// 		// 执行工作流
// 		web.RunHandlers(ctx, g.handlers)
// 		// 执行终端函数
// 		handler(ctx)
// 	})
// }

func (g *GinServer) AddNode(handlers ...web.NodeHandler) {
	g.handlers = append(g.handlers, handlers...)
}

func (g *GinServer) Use(middleware ...gin.HandlerFunc) {
	g.Engine.Use(middleware...)
}

func (g *GinServer) GET(path string, handler web.RequestHandler) {
	g.Engine.GET(path, func(context *gin.Context) {
		handler(context)
	})
}

func (g *GinServer) POST(path string, handler web.RequestHandler) {
	g.Engine.POST(path, func(context *gin.Context) {
		handler(context)
	})
}

func (g *GinServer) Start() {
	go func() {
		g.Engine.Run(fmt.Sprintf(":%d", g.config.Port))
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
