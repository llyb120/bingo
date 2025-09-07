package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Set a key/value
func Set[T any](conn *redis.Client, key string, data T, t time.Duration) error {
	if t == 0 {
		t = 10 * time.Minute
	}
	var (
		reply string
		err   error
	)
	// now := time.Now()
	// span := skyWalking.SkyWalkingStartSpanForRedis(c, "SET:"+key, "")
	// defer func() {
	// 	skyWalking.EndSpanForRedisCache(c, span, "SET", reply, time.Now().Sub(now), ext, err)
	// }()

	var value []byte
	switch data := any(data).(type) {
	case []byte:
		value = data
	default:
		var err error
		value, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}
	ctx := context.Background()
	reply, err = conn.Set(ctx, key, string(value), t).Result()
	if err != nil {
		return err
	}

	_ = reply

	return nil
}

// func SetCache(c *gin.Context, conn *redis.Client, key string, data []byte, t time.Duration, ext string) error {
// 	var (
// 		// reply string
// 		err error
// 	)
// 	// now := time.Now()
// 	// span := skyWalking.SkyWalkingStartSpanForRedis(c, "SET:"+key, "")
// 	// defer func() {
// 	// 	skyWalking.EndSpanForRedisCache(c, span, "SET", reply, time.Now().Sub(now), ext, err)
// 	// }()

// 	//value, err := sonics.Marshal(data)
// 	//if err != nil {
// 	//	return err
// 	//}

// 	ctx := context.Background()
// 	// zero copy
// 	reply, err = conn.Set(ctx, key, data, t).Result()
// 	if err != nil {
// 		return err
// 	}

// 	_ = reply

// 	return nil
// }

// Get get a key
func Get(conn *redis.Client, key string) ([]byte, error) {
	var (
		reply     string
		gerr, err error
	)
	// now := time.Now()
	// span := skyWalking.SkyWalkingStartSpanForRedis(c, "GET:"+key, "")
	// defer func() {
	// 	redisRes := reply
	// 	skyWalking.EndSpanForRedisCache(c, span, "GET", redisRes[:utils.Min(len(redisRes), 1000)], time.Now().Sub(now), ext, gerr)
	// }()

	ctx := context.Background()
	reply, err = conn.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		gerr = err
		// logging.ErrorF(c, "redis get key:%s err:%s", key, err)
		return nil, gerr
	}

	return []byte(reply), nil
}
