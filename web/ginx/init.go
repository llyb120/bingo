package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
)

// var globalGin *gin.Engine
// var globalState *core.State

var GinStarter core.Starter = func() func() {
	// globalState = state
	// globalGin = r
	r := gin.New()
	var server = ginServer{
		Port: 8080,
	}
	config := core.Require[config.Config]()
	config.LoadToStruct("server", &server)

	gins := &GinServer{Engine: r, config: server}
	r.Use(func(c *gin.Context) {
		c.Set("$bingo-ginx-handlers", gins.handlers)
	})

	core.ExportInstance(gins, core.RegisterOption{Name: "Gin"})
	return nil
}

type ginServer struct {
	Port int `json:"port"`
}
