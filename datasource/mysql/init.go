package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/yoya2/y"
)

// func init() {
// 	// 读取配置文件，初始化db
// 	core.On(config.EVENT_LOAD_CONFIG_OK, func() {
// 		core.RegisterInstance("mysql-db", &Mysql{})
// 	})
// }

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var MysqlStarter core.Starter = func() func() {
	var cfg config.Config
	core.Require(&cfg)
	var cfgs map[string]MysqlConfig
	cfg.LoadToStruct("datasource.mysql", &cfgs)

	// 初始化数据源
	dbs := y.Flex(y.Keys(cfgs), func(name string, _ int) *sql.DB {
		cfg := cfgs[name]
		db, err := openMysqlConnection(cfg)
		if err != nil {
			panic(err)
		}
		if err := db.Ping(); err != nil {
			panic(err)
		}
		core.ExportInstance(db, core.RegisterOption{Name: name})
		return db
	}, y.UseAsync, y.UsePanic)

	return func() {
		y.Flex(dbs, func(db *sql.DB, _ int) any {
			db.Close()
			return nil
		}, y.UseAsync)
	}
}

func openMysqlConnection(cfg MysqlConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
