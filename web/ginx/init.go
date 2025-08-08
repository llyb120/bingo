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

var GinStarter core.Starter = &ginStarter{}

type ginStarter struct {
}

func (g *ginStarter) Init(state *core.State) {
	// globalState = state
	// globalGin = r
	r := gin.New()
	core.ExportInstance(state, r, core.RegisterOption{Name: "Gin"})
}

func (g *ginStarter) Destroy(state *core.State) {

}

type ginServer struct {
	Port int `json:"port"`
}

func Start(state *core.State) {
	var gin *gin.Engine
	var config config.Config
	var server = ginServer{
		Port: 8080,
	}
	state.Use(&gin)
	state.Use(&config)
	config.LoadToStruct("server", &server)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", server.Port),
		Handler: gin,
	}
	log.Printf("[info] start http server listening %s", server.Port)

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
