package ginx

import (
	"fmt"

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
	g.Run(fmt.Sprintf(":%d", g.config.Port))
}
