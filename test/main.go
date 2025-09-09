package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/datasource/mysql"
	"github.com/llyb120/bingo/datasource/redis"
	"github.com/llyb120/bingo/sqly/tool"
	"github.com/llyb120/bingo/test/sql/template"
	"github.com/llyb120/bingo/web/ginx"
)

var (
	mysql0 = core.Use[sql.DB]("m0")
	mysql1 = core.Use[sql.DB]()
)

var plugins = []core.Starter{
	config.ConfigStarter,
	mysql.MysqlStarter,
	ginx.GinStarter,
	redis.RedisStarter,
	RouterStarter,
	template.TestTemplateStarter,
	tool.GoTemplateStarter,
	//cache.CacheStarter,
}

func init() {
	os.Setenv("BINGO_CONFIG_PATH", "./config.properties")
}

func main() {
	core.Boot(plugins...)
	defer core.Shutdown()

	// core.Use(&mysql0, "m0")
	// core.Use(&mysql1)

	var gin = core.Require[ginx.GinServer]()
	defer gin.Start()
}

type TestValidate struct {
	TopSize string `json:"top_size"`
}

func (t *TestValidate) Validate() error {
	if t.TopSize == "" {
		return fmt.Errorf("top_size is empty")
	}
	return nil
}
