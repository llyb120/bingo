package config

import (
	"os"

	"github.com/llyb120/bingo/core"
)

var (
	ConfigStarter core.Starter = func() func() {
		cfg := LoadConfig()
		core.ExportInstance(cfg, core.RegisterOption{Name: "config"})

		return nil
	}
)

// LoadConfig 加载配置文件
func LoadConfig() Config {
	path := os.Getenv("BINGO_CONFIG_PATH")
	if path == "" {
		path = "./config.properties"
	}
	return mustLoadProperties(path)
}
