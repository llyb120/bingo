package ginx

import (
	"fmt"
	"log"
	"net/http"

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
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", g.config.Port),
		Handler: g.Engine,
	}
	log.Printf("[info] start http server listening %d", g.config.Port)

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
