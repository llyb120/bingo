package cache

import (
	"time"
)

// Func_0_1: 0个参数，1个返回值
func Func_0_1[R0 any](fn func() R0, keyGenerator func() string, ttlFn func() time.Duration) func() R0 {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func() (r0 R0) {
		key := prefix + ":" + keyGenerator()
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			return r0
		}
		r0 = fn()
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0
		}
		defer Set(redisConn, key, []any{r0}, ttlFn())
		return r0
	}
}

// Func_0_2: 0个参数，2个返回值
func Func_0_2[R0, R1 any](fn func() (R0, R1), keyGenerator func() string, ttlFn func() time.Duration) func() (R0, R1) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func() (r0 R0, r1 R1) {
		key := prefix + ":" + keyGenerator()
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
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
		r0, r1 = fn()
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1
		}
		defer Set(redisConn, key, []any{r0, r1}, ttlFn())
		return r0, r1
	}
}

// Func_0_3: 0个参数，3个返回值
func Func_0_3[R0, R1, R2 any](fn func() (R0, R1, R2), keyGenerator func() string, ttlFn func() time.Duration) func() (R0, R1, R2) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func() (r0 R0, r1 R1, r2 R2) {
		key := prefix + ":" + keyGenerator()
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			return r0, r1, r2
		}
		r0, r1, r2 = fn()
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2
		}
		defer Set(redisConn, key, []any{r0, r1, r2}, ttlFn())
		return r0, r1, r2
	}
}

// Func_0_4: 0个参数，4个返回值
func Func_0_4[R0, R1, R2, R3 any](fn func() (R0, R1, R2, R3), keyGenerator func() string, ttlFn func() time.Duration) func() (R0, R1, R2, R3) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func() (r0 R0, r1 R1, r2 R2, r3 R3) {
		key := prefix + ":" + keyGenerator()
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			return r0, r1, r2, r3
		}
		r0, r1, r2, r3 = fn()
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3}, ttlFn())
		return r0, r1, r2, r3
	}
}

// Func_0_5: 0个参数，5个返回值
func Func_0_5[R0, R1, R2, R3, R4 any](fn func() (R0, R1, R2, R3, R4), keyGenerator func() string, ttlFn func() time.Duration) func() (R0, R1, R2, R3, R4) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func() (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4) {
		key := prefix + ":" + keyGenerator()
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			return r0, r1, r2, r3, r4
		}
		r0, r1, r2, r3, r4 = fn()
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4}, ttlFn())
		return r0, r1, r2, r3, r4
	}
}

// Func_0_6: 0个参数，6个返回值
func Func_0_6[R0, R1, R2, R3, R4, R5 any](fn func() (R0, R1, R2, R3, R4, R5), keyGenerator func() string, ttlFn func() time.Duration) func() (R0, R1, R2, R3, R4, R5) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func() (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		key := prefix + ":" + keyGenerator()
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4, r5
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			if len(res) > 5 {
				r5, _ = res[5].(R5)
			}
			return r0, r1, r2, r3, r4, r5
		}
		r0, r1, r2, r3, r4, r5 = fn()
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r5).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4, r5}, ttlFn())
		return r0, r1, r2, r3, r4, r5
	}
}

// Func_1_1: 1个参数，1个返回值
func Func_1_1[P0, R0 any](fn func(P0) R0, keyGenerator func(P0) string, ttlFn func(P0) time.Duration) func(P0) R0 {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0) (r0 R0) {
		key := prefix + ":" + keyGenerator(p0)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			return r0
		}
		r0 = fn(p0)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0
		}
		defer Set(redisConn, key, []any{r0}, ttlFn(p0))
		return r0
	}
}

// Func_1_2: 1个参数，2个返回值
func Func_1_2[P0, R0, R1 any](fn func(P0) (R0, R1), keyGenerator func(P0) string, ttlFn func(P0) time.Duration) func(P0) (R0, R1) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0) (r0 R0, r1 R1) {
		key := prefix + ":" + keyGenerator(p0)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
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
		r0, r1 = fn(p0)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1
		}
		defer Set(redisConn, key, []any{r0, r1}, ttlFn(p0))
		return r0, r1
	}
}

// Func_1_3: 1个参数，3个返回值
func Func_1_3[P0, R0, R1, R2 any](fn func(P0) (R0, R1, R2), keyGenerator func(P0) string, ttlFn func(P0) time.Duration) func(P0) (R0, R1, R2) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0) (r0 R0, r1 R1, r2 R2) {
		key := prefix + ":" + keyGenerator(p0)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			return r0, r1, r2
		}
		r0, r1, r2 = fn(p0)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2
		}
		defer Set(redisConn, key, []any{r0, r1, r2}, ttlFn(p0))
		return r0, r1, r2
	}
}

// Func_1_4: 1个参数，4个返回值
func Func_1_4[P0, R0, R1, R2, R3 any](fn func(P0) (R0, R1, R2, R3), keyGenerator func(P0) string, ttlFn func(P0) time.Duration) func(P0) (R0, R1, R2, R3) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0) (r0 R0, r1 R1, r2 R2, r3 R3) {
		key := prefix + ":" + keyGenerator(p0)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			return r0, r1, r2, r3
		}
		r0, r1, r2, r3 = fn(p0)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3}, ttlFn(p0))
		return r0, r1, r2, r3
	}
}

// Func_1_5: 1个参数，5个返回值
func Func_1_5[P0, R0, R1, R2, R3, R4 any](fn func(P0) (R0, R1, R2, R3, R4), keyGenerator func(P0) string, ttlFn func(P0) time.Duration) func(P0) (R0, R1, R2, R3, R4) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4) {
		key := prefix + ":" + keyGenerator(p0)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			return r0, r1, r2, r3, r4
		}
		r0, r1, r2, r3, r4 = fn(p0)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4}, ttlFn(p0))
		return r0, r1, r2, r3, r4
	}
}

// Func_1_6: 1个参数，6个返回值
func Func_1_6[P0, R0, R1, R2, R3, R4, R5 any](fn func(P0) (R0, R1, R2, R3, R4, R5), keyGenerator func(P0) string, ttlFn func(P0) time.Duration) func(P0) (R0, R1, R2, R3, R4, R5) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		key := prefix + ":" + keyGenerator(p0)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4, r5
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			if len(res) > 5 {
				r5, _ = res[5].(R5)
			}
			return r0, r1, r2, r3, r4, r5
		}
		r0, r1, r2, r3, r4, r5 = fn(p0)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r5).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4, r5}, ttlFn(p0))
		return r0, r1, r2, r3, r4, r5
	}
}

// Func_2_1: 2个参数，1个返回值
func Func_2_1[P0, P1, R0 any](fn func(P0, P1) R0, keyGenerator func(P0, P1) string, ttlFn func(P0, P1) time.Duration) func(P0, P1) R0 {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1) (r0 R0) {
		key := prefix + ":" + keyGenerator(p0, p1)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			return r0
		}
		r0 = fn(p0, p1)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0
		}
		defer Set(redisConn, key, []any{r0}, ttlFn(p0, p1))
		return r0
	}
}

// Func_2_3: 2个参数，3个返回值
func Func_2_3[P0, P1, R0, R1, R2 any](fn func(P0, P1) (R0, R1, R2), keyGenerator func(P0, P1) string, ttlFn func(P0, P1) time.Duration) func(P0, P1) (R0, R1, R2) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1) (r0 R0, r1 R1, r2 R2) {
		key := prefix + ":" + keyGenerator(p0, p1)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			return r0, r1, r2
		}
		r0, r1, r2 = fn(p0, p1)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2
		}
		defer Set(redisConn, key, []any{r0, r1, r2}, ttlFn(p0, p1))
		return r0, r1, r2
	}
}

// Func_2_4: 2个参数，4个返回值
func Func_2_4[P0, P1, R0, R1, R2, R3 any](fn func(P0, P1) (R0, R1, R2, R3), keyGenerator func(P0, P1) string, ttlFn func(P0, P1) time.Duration) func(P0, P1) (R0, R1, R2, R3) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1) (r0 R0, r1 R1, r2 R2, r3 R3) {
		key := prefix + ":" + keyGenerator(p0, p1)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			return r0, r1, r2, r3
		}
		r0, r1, r2, r3 = fn(p0, p1)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3}, ttlFn(p0, p1))
		return r0, r1, r2, r3
	}
}

// Func_2_5: 2个参数，5个返回值
func Func_2_5[P0, P1, R0, R1, R2, R3, R4 any](fn func(P0, P1) (R0, R1, R2, R3, R4), keyGenerator func(P0, P1) string, ttlFn func(P0, P1) time.Duration) func(P0, P1) (R0, R1, R2, R3, R4) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4) {
		key := prefix + ":" + keyGenerator(p0, p1)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			return r0, r1, r2, r3, r4
		}
		r0, r1, r2, r3, r4 = fn(p0, p1)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4}, ttlFn(p0, p1))
		return r0, r1, r2, r3, r4
	}
}

// Func_2_6: 2个参数，6个返回值
func Func_2_6[P0, P1, R0, R1, R2, R3, R4, R5 any](fn func(P0, P1) (R0, R1, R2, R3, R4, R5), keyGenerator func(P0, P1) string, ttlFn func(P0, P1) time.Duration) func(P0, P1) (R0, R1, R2, R3, R4, R5) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		key := prefix + ":" + keyGenerator(p0, p1)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4, r5
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			if len(res) > 5 {
				r5, _ = res[5].(R5)
			}
			return r0, r1, r2, r3, r4, r5
		}
		r0, r1, r2, r3, r4, r5 = fn(p0, p1)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r5).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4, r5}, ttlFn(p0, p1))
		return r0, r1, r2, r3, r4, r5
	}
}

// Func_3_1: 3个参数，1个返回值
func Func_3_1[P0, P1, P2, R0 any](fn func(P0, P1, P2) R0, keyGenerator func(P0, P1, P2) string, ttlFn func(P0, P1, P2) time.Duration) func(P0, P1, P2) R0 {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2) (r0 R0) {
		key := prefix + ":" + keyGenerator(p0, p1, p2)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			return r0
		}
		r0 = fn(p0, p1, p2)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0
		}
		defer Set(redisConn, key, []any{r0}, ttlFn(p0, p1, p2))
		return r0
	}
}

// Func_3_2: 3个参数，2个返回值
func Func_3_2[P0, P1, P2, R0, R1 any](fn func(P0, P1, P2) (R0, R1), keyGenerator func(P0, P1, P2) string, ttlFn func(P0, P1, P2) time.Duration) func(P0, P1, P2) (R0, R1) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2) (r0 R0, r1 R1) {
		key := prefix + ":" + keyGenerator(p0, p1, p2)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
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
		r0, r1 = fn(p0, p1, p2)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1
		}
		defer Set(redisConn, key, []any{r0, r1}, ttlFn(p0, p1, p2))
		return r0, r1
	}
}

// Func_3_3: 3个参数，3个返回值
func Func_3_3[P0, P1, P2, R0, R1, R2 any](fn func(P0, P1, P2) (R0, R1, R2), keyGenerator func(P0, P1, P2) string, ttlFn func(P0, P1, P2) time.Duration) func(P0, P1, P2) (R0, R1, R2) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2) (r0 R0, r1 R1, r2 R2) {
		key := prefix + ":" + keyGenerator(p0, p1, p2)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			return r0, r1, r2
		}
		r0, r1, r2 = fn(p0, p1, p2)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2
		}
		defer Set(redisConn, key, []any{r0, r1, r2}, ttlFn(p0, p1, p2))
		return r0, r1, r2
	}
}

// Func_3_4: 3个参数，4个返回值
func Func_3_4[P0, P1, P2, R0, R1, R2, R3 any](fn func(P0, P1, P2) (R0, R1, R2, R3), keyGenerator func(P0, P1, P2) string, ttlFn func(P0, P1, P2) time.Duration) func(P0, P1, P2) (R0, R1, R2, R3) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2) (r0 R0, r1 R1, r2 R2, r3 R3) {
		key := prefix + ":" + keyGenerator(p0, p1, p2)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			return r0, r1, r2, r3
		}
		r0, r1, r2, r3 = fn(p0, p1, p2)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3}, ttlFn(p0, p1, p2))
		return r0, r1, r2, r3
	}
}

// Func_3_5: 3个参数，5个返回值
func Func_3_5[P0, P1, P2, R0, R1, R2, R3, R4 any](fn func(P0, P1, P2) (R0, R1, R2, R3, R4), keyGenerator func(P0, P1, P2) string, ttlFn func(P0, P1, P2) time.Duration) func(P0, P1, P2) (R0, R1, R2, R3, R4) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4) {
		key := prefix + ":" + keyGenerator(p0, p1, p2)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			return r0, r1, r2, r3, r4
		}
		r0, r1, r2, r3, r4 = fn(p0, p1, p2)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4}, ttlFn(p0, p1, p2))
		return r0, r1, r2, r3, r4
	}
}

// Func_3_6: 3个参数，6个返回值
func Func_3_6[P0, P1, P2, R0, R1, R2, R3, R4, R5 any](fn func(P0, P1, P2) (R0, R1, R2, R3, R4, R5), keyGenerator func(P0, P1, P2) string, ttlFn func(P0, P1, P2) time.Duration) func(P0, P1, P2) (R0, R1, R2, R3, R4, R5) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		key := prefix + ":" + keyGenerator(p0, p1, p2)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4, r5
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			if len(res) > 5 {
				r5, _ = res[5].(R5)
			}
			return r0, r1, r2, r3, r4, r5
		}
		r0, r1, r2, r3, r4, r5 = fn(p0, p1, p2)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r5).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4, r5}, ttlFn(p0, p1, p2))
		return r0, r1, r2, r3, r4, r5
	}
}

// Func_4_1: 4个参数，1个返回值
func Func_4_1[P0, P1, P2, P3, R0 any](fn func(P0, P1, P2, P3) R0, keyGenerator func(P0, P1, P2, P3) string, ttlFn func(P0, P1, P2, P3) time.Duration) func(P0, P1, P2, P3) R0 {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (r0 R0) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			return r0
		}
		r0 = fn(p0, p1, p2, p3)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0
		}
		defer Set(redisConn, key, []any{r0}, ttlFn(p0, p1, p2, p3))
		return r0
	}
}

// Func_4_2: 4个参数，2个返回值
func Func_4_2[P0, P1, P2, P3, R0, R1 any](fn func(P0, P1, P2, P3) (R0, R1), keyGenerator func(P0, P1, P2, P3) string, ttlFn func(P0, P1, P2, P3) time.Duration) func(P0, P1, P2, P3) (R0, R1) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (r0 R0, r1 R1) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
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
		r0, r1 = fn(p0, p1, p2, p3)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1
		}
		defer Set(redisConn, key, []any{r0, r1}, ttlFn(p0, p1, p2, p3))
		return r0, r1
	}
}

// Func_4_3: 4个参数，3个返回值
func Func_4_3[P0, P1, P2, P3, R0, R1, R2 any](fn func(P0, P1, P2, P3) (R0, R1, R2), keyGenerator func(P0, P1, P2, P3) string, ttlFn func(P0, P1, P2, P3) time.Duration) func(P0, P1, P2, P3) (R0, R1, R2) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (r0 R0, r1 R1, r2 R2) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			return r0, r1, r2
		}
		r0, r1, r2 = fn(p0, p1, p2, p3)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2
		}
		defer Set(redisConn, key, []any{r0, r1, r2}, ttlFn(p0, p1, p2, p3))
		return r0, r1, r2
	}
}

// Func_4_4: 4个参数，4个返回值
func Func_4_4[P0, P1, P2, P3, R0, R1, R2, R3 any](fn func(P0, P1, P2, P3) (R0, R1, R2, R3), keyGenerator func(P0, P1, P2, P3) string, ttlFn func(P0, P1, P2, P3) time.Duration) func(P0, P1, P2, P3) (R0, R1, R2, R3) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (r0 R0, r1 R1, r2 R2, r3 R3) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			return r0, r1, r2, r3
		}
		r0, r1, r2, r3 = fn(p0, p1, p2, p3)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3}, ttlFn(p0, p1, p2, p3))
		return r0, r1, r2, r3
	}
}

// Func_4_5: 4个参数，5个返回值
func Func_4_5[P0, P1, P2, P3, R0, R1, R2, R3, R4 any](fn func(P0, P1, P2, P3) (R0, R1, R2, R3, R4), keyGenerator func(P0, P1, P2, P3) string, ttlFn func(P0, P1, P2, P3) time.Duration) func(P0, P1, P2, P3) (R0, R1, R2, R3, R4) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			return r0, r1, r2, r3, r4
		}
		r0, r1, r2, r3, r4 = fn(p0, p1, p2, p3)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4}, ttlFn(p0, p1, p2, p3))
		return r0, r1, r2, r3, r4
	}
}

// Func_4_6: 4个参数，6个返回值
func Func_4_6[P0, P1, P2, P3, R0, R1, R2, R3, R4, R5 any](fn func(P0, P1, P2, P3) (R0, R1, R2, R3, R4, R5), keyGenerator func(P0, P1, P2, P3) string, ttlFn func(P0, P1, P2, P3) time.Duration) func(P0, P1, P2, P3) (R0, R1, R2, R3, R4, R5) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4, r5
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			if len(res) > 5 {
				r5, _ = res[5].(R5)
			}
			return r0, r1, r2, r3, r4, r5
		}
		r0, r1, r2, r3, r4, r5 = fn(p0, p1, p2, p3)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r5).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4, r5}, ttlFn(p0, p1, p2, p3))
		return r0, r1, r2, r3, r4, r5
	}
}

// Func_5_1: 5个参数，1个返回值
func Func_5_1[P0, P1, P2, P3, P4, R0 any](fn func(P0, P1, P2, P3, P4) R0, keyGenerator func(P0, P1, P2, P3, P4) string, ttlFn func(P0, P1, P2, P3, P4) time.Duration) func(P0, P1, P2, P3, P4) R0 {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (r0 R0) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			return r0
		}
		r0 = fn(p0, p1, p2, p3, p4)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0
		}
		defer Set(redisConn, key, []any{r0}, ttlFn(p0, p1, p2, p3, p4))
		return r0
	}
}

// Func_5_2: 5个参数，2个返回值
func Func_5_2[P0, P1, P2, P3, P4, R0, R1 any](fn func(P0, P1, P2, P3, P4) (R0, R1), keyGenerator func(P0, P1, P2, P3, P4) string, ttlFn func(P0, P1, P2, P3, P4) time.Duration) func(P0, P1, P2, P3, P4) (R0, R1) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (r0 R0, r1 R1) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
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
		r0, r1 = fn(p0, p1, p2, p3, p4)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1
		}
		defer Set(redisConn, key, []any{r0, r1}, ttlFn(p0, p1, p2, p3, p4))
		return r0, r1
	}
}

// Func_5_3: 5个参数，3个返回值
func Func_5_3[P0, P1, P2, P3, P4, R0, R1, R2 any](fn func(P0, P1, P2, P3, P4) (R0, R1, R2), keyGenerator func(P0, P1, P2, P3, P4) string, ttlFn func(P0, P1, P2, P3, P4) time.Duration) func(P0, P1, P2, P3, P4) (R0, R1, R2) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (r0 R0, r1 R1, r2 R2) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			return r0, r1, r2
		}
		r0, r1, r2 = fn(p0, p1, p2, p3, p4)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2
		}
		defer Set(redisConn, key, []any{r0, r1, r2}, ttlFn(p0, p1, p2, p3, p4))
		return r0, r1, r2
	}
}

// Func_5_4: 5个参数，4个返回值
func Func_5_4[P0, P1, P2, P3, P4, R0, R1, R2, R3 any](fn func(P0, P1, P2, P3, P4) (R0, R1, R2, R3), keyGenerator func(P0, P1, P2, P3, P4) string, ttlFn func(P0, P1, P2, P3, P4) time.Duration) func(P0, P1, P2, P3, P4) (R0, R1, R2, R3) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (r0 R0, r1 R1, r2 R2, r3 R3) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			return r0, r1, r2, r3
		}
		r0, r1, r2, r3 = fn(p0, p1, p2, p3, p4)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3}, ttlFn(p0, p1, p2, p3, p4))
		return r0, r1, r2, r3
	}
}

// Func_5_5: 5个参数，5个返回值
func Func_5_5[P0, P1, P2, P3, P4, R0, R1, R2, R3, R4 any](fn func(P0, P1, P2, P3, P4) (R0, R1, R2, R3, R4), keyGenerator func(P0, P1, P2, P3, P4) string, ttlFn func(P0, P1, P2, P3, P4) time.Duration) func(P0, P1, P2, P3, P4) (R0, R1, R2, R3, R4) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			return r0, r1, r2, r3, r4
		}
		r0, r1, r2, r3, r4 = fn(p0, p1, p2, p3, p4)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4}, ttlFn(p0, p1, p2, p3, p4))
		return r0, r1, r2, r3, r4
	}
}

// Func_5_6: 5个参数，6个返回值
func Func_5_6[P0, P1, P2, P3, P4, R0, R1, R2, R3, R4, R5 any](fn func(P0, P1, P2, P3, P4) (R0, R1, R2, R3, R4, R5), keyGenerator func(P0, P1, P2, P3, P4) string, ttlFn func(P0, P1, P2, P3, P4) time.Duration) func(P0, P1, P2, P3, P4) (R0, R1, R2, R3, R4, R5) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4, r5
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			if len(res) > 5 {
				r5, _ = res[5].(R5)
			}
			return r0, r1, r2, r3, r4, r5
		}
		r0, r1, r2, r3, r4, r5 = fn(p0, p1, p2, p3, p4)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r5).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4, r5}, ttlFn(p0, p1, p2, p3, p4))
		return r0, r1, r2, r3, r4, r5
	}
}

// Func_6_1: 6个参数，1个返回值
func Func_6_1[P0, P1, P2, P3, P4, P5, R0 any](fn func(P0, P1, P2, P3, P4, P5) R0, keyGenerator func(P0, P1, P2, P3, P4, P5) string, ttlFn func(P0, P1, P2, P3, P4, P5) time.Duration) func(P0, P1, P2, P3, P4, P5) R0 {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (r0 R0) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4, p5)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			return r0
		}
		r0 = fn(p0, p1, p2, p3, p4, p5)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0
		}
		defer Set(redisConn, key, []any{r0}, ttlFn(p0, p1, p2, p3, p4, p5))
		return r0
	}
}

// Func_6_2: 6个参数，2个返回值
func Func_6_2[P0, P1, P2, P3, P4, P5, R0, R1 any](fn func(P0, P1, P2, P3, P4, P5) (R0, R1), keyGenerator func(P0, P1, P2, P3, P4, P5) string, ttlFn func(P0, P1, P2, P3, P4, P5) time.Duration) func(P0, P1, P2, P3, P4, P5) (R0, R1) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (r0 R0, r1 R1) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4, p5)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
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
		r0, r1 = fn(p0, p1, p2, p3, p4, p5)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1
		}
		defer Set(redisConn, key, []any{r0, r1}, ttlFn(p0, p1, p2, p3, p4, p5))
		return r0, r1
	}
}

// Func_6_3: 6个参数，3个返回值
func Func_6_3[P0, P1, P2, P3, P4, P5, R0, R1, R2 any](fn func(P0, P1, P2, P3, P4, P5) (R0, R1, R2), keyGenerator func(P0, P1, P2, P3, P4, P5) string, ttlFn func(P0, P1, P2, P3, P4, P5) time.Duration) func(P0, P1, P2, P3, P4, P5) (R0, R1, R2) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (r0 R0, r1 R1, r2 R2) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4, p5)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			return r0, r1, r2
		}
		r0, r1, r2 = fn(p0, p1, p2, p3, p4, p5)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2
		}
		defer Set(redisConn, key, []any{r0, r1, r2}, ttlFn(p0, p1, p2, p3, p4, p5))
		return r0, r1, r2
	}
}

// Func_6_4: 6个参数，4个返回值
func Func_6_4[P0, P1, P2, P3, P4, P5, R0, R1, R2, R3 any](fn func(P0, P1, P2, P3, P4, P5) (R0, R1, R2, R3), keyGenerator func(P0, P1, P2, P3, P4, P5) string, ttlFn func(P0, P1, P2, P3, P4, P5) time.Duration) func(P0, P1, P2, P3, P4, P5) (R0, R1, R2, R3) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (r0 R0, r1 R1, r2 R2, r3 R3) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4, p5)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			return r0, r1, r2, r3
		}
		r0, r1, r2, r3 = fn(p0, p1, p2, p3, p4, p5)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3}, ttlFn(p0, p1, p2, p3, p4, p5))
		return r0, r1, r2, r3
	}
}

// Func_6_5: 6个参数，5个返回值
func Func_6_5[P0, P1, P2, P3, P4, P5, R0, R1, R2, R3, R4 any](fn func(P0, P1, P2, P3, P4, P5) (R0, R1, R2, R3, R4), keyGenerator func(P0, P1, P2, P3, P4, P5) string, ttlFn func(P0, P1, P2, P3, P4, P5) time.Duration) func(P0, P1, P2, P3, P4, P5) (R0, R1, R2, R3, R4) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4, p5)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			return r0, r1, r2, r3, r4
		}
		r0, r1, r2, r3, r4 = fn(p0, p1, p2, p3, p4, p5)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4}, ttlFn(p0, p1, p2, p3, p4, p5))
		return r0, r1, r2, r3, r4
	}
}

// Func_6_6: 6个参数，6个返回值
func Func_6_6[P0, P1, P2, P3, P4, P5, R0, R1, R2, R3, R4, R5 any](fn func(P0, P1, P2, P3, P4, P5) (R0, R1, R2, R3, R4, R5), keyGenerator func(P0, P1, P2, P3, P4, P5) string, ttlFn func(P0, P1, P2, P3, P4, P5) time.Duration) func(P0, P1, P2, P3, P4, P5) (R0, R1, R2, R3, R4, R5) {
	prefix := cfg.GetString("cache.prefix") + ":" + cfg.GetString("server.environment")
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (r0 R0, r1 R1, r2 R2, r3 R3, r4 R4, r5 R5) {
		key := prefix + ":" + keyGenerator(p0, p1, p2, p3, p4, p5)
		if bs, err := Get(redisConn, key); err == nil && len(bs) > 0 {
			var res = []any{}
			err := json.Unmarshal(bs, &res)
			if err != nil {
				return r0, r1, r2, r3, r4, r5
			}
			if len(res) > 0 {
				r0, _ = res[0].(R0)
			}
			if len(res) > 1 {
				r1, _ = res[1].(R1)
			}
			if len(res) > 2 {
				r2, _ = res[2].(R2)
			}
			if len(res) > 3 {
				r3, _ = res[3].(R3)
			}
			if len(res) > 4 {
				r4, _ = res[4].(R4)
			}
			if len(res) > 5 {
				r5, _ = res[5].(R5)
			}
			return r0, r1, r2, r3, r4, r5
		}
		r0, r1, r2, r3, r4, r5 = fn(p0, p1, p2, p3, p4, p5)
		// 检查是否有error，如果有，则不设置缓存
		if _, ok := any(r0).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r1).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r2).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r3).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r4).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		if _, ok := any(r5).(error); ok {
			return r0, r1, r2, r3, r4, r5
		}
		defer Set(redisConn, key, []any{r0, r1, r2, r3, r4, r5}, ttlFn(p0, p1, p2, p3, p4, p5))
		return r0, r1, r2, r3, r4, r5
	}
}

