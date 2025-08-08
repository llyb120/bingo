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

var GinStarter core.Starter = func() func() {
	// globalState = state
	// globalGin = r
	r := gin.New()
	core.ExportInstance(r, core.RegisterOption{Name: "Gin"})

	return nil
}

type ginServer struct {
	Port int `json:"port"`
}

func Start() {
	var gin *gin.Engine
	var config config.Config
	var server = ginServer{
		Port: 8080,
	}
	core.Use(&gin)
	core.Use(&config)
	config.LoadToStruct("server", &server)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", server.Port),
		Handler: gin,
	}
	log.Printf("[info] start http server listening %d", server.Port)

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
