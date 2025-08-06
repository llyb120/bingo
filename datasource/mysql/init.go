package mysql

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
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

var MysqlStarter core.Starter = &mysqlStarter{}

type mysqlStarter struct {
}

func (m *mysqlStarter) Init(state *core.State) {
	var cfg config.Config
	state.Require(&cfg)
	var cfgs map[string]MysqlConfig
	cfg.LoadToStruct("datasource.mysql", &cfgs)

	var g sync.WaitGroup
	for name, cfg := range cfgs {
		g.Add(1)
		go func() {
			defer g.Done()
			// 连接
			db, err := openMysqlConnection(cfg)
			if err != nil {
				panic(err)
			}
			if db == nil {
				panic("failed to open mysql connection")
			}
			// 注册
			state.ExportInstance(db, core.RegisterOption{Name: name})
		}()
	}
	g.Wait()
}

func (m *mysqlStarter) Destroy(state *core.State) {

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
