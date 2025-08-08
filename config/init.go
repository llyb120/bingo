package config

import (
	"os"

	"github.com/llyb120/bingo/core"
)

var (
	ConfigStarter core.Starter = &configStarter{}
)

type configStarter struct {
}

func (c *configStarter) Init(state *core.State) {
	cfg := LoadConfig()
	core.ExportInstance(state, cfg, core.RegisterOption{Name: "config"})
}

func (c *configStarter) Destroy(state *core.State) {

}

// LoadConfig 加载配置文件
func LoadConfig() Config {
	path := os.Getenv("BINGO_CONFIG_PATH")
	if path == "" {
		path = "./config.properties"
	}
	return mustLoadProperties(path)
}
