//go:generate go run generate_funcs.go
package cache

import (
	"time"

	"github.com/llyb120/bingo/config"

	"github.com/go-redis/redis/v8"
	"github.com/llyb120/bingo/core"
)

var redisConn = core.Use[*redis.Client]()
var cfg = core.Use[*config.Config]("config")

func Func_2_2[P0, P1, R0, R1 any](fn func(P0, P1) (R0, R1), keyGenerator func(P0, P1) string, ttlFn func(P0, P1) time.Duration) func(P0, P1) (R0, R1) {
	prefix := cfg().GetString("cache.prefix") + ":" + cfg().GetString("server.environment")
	return func(p0 P0, p1 P1) (r0 R0, r1 R1) {
		key := prefix + ":" + keyGenerator(p0, p1)
		if bs, err := Get(redisConn(), key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			return r0, r1
		}
		r0, r1 = fn(p0, p1)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1
		}
		defer Set(redisConn(), key, []any{r0, r1}, ttlFn(p0, p1))
		return r0, r1
	}
}
