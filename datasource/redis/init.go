package redis

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/llyb120/bingo/config"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/yoya2/y"
)

var RedisStarter core.Starter = func() func() {
	// require
	var cfg config.Config
	core.Require(&cfg)
	var cfgs map[string]redisConfig
	cfg.LoadToStruct("datasource.redis", &cfgs)

	dbs := y.Flex(y.Keys(cfgs), func(name string, _ int) *redis.Client {
		v := cfgs[name]
		if !v.Enable {
			return nil
		}
		port := strconv.FormatInt(int64(v.Port), 10)
		addr := v.Host + ":" + port
		db := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: v.Password,
			DB:       0,
			// PoolSize: 50,
		})
		if err := db.Ping(context.Background()).Err(); err != nil {
			panic(err)
		}
		core.ExportInstance(db, core.RegisterOption{Name: name})
		return db
	}, y.UseAsync, y.UsePanic, y.NotNil)

	// 清理函数
	return func() {
		y.Flex(dbs, func(db *redis.Client, _ int) any {
			db.Close()
			return nil
		}, y.UseAsync)
	}
}

type redisConfig struct {
	Enable   bool   `json:"enable"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}
