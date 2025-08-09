package cache

import (
	"fmt"

	"github.com/llyb120/bingo/core"
)

var CacheStarter core.Starter = func() func() {

	// 使用事件订阅机制对命名导出函数进行劫持/增强
	core.On("before:sql.Select", func(args ...any) {
		call := args[0].(*core.FuncCall)
		// 示例：可根据 key 做缓存命中
		fmt.Println("intercept sql.Select")
		call.Skip = true
		// 如果命中，可设置 call.Result 并 call.Skip = true 短路原调用
		_ = call
	})

	core.On("after:sql.Select", func(args ...any) {
		call := args[0].(*core.FuncCall)
		// 示例：将返回结果写入缓存
		_ = call
	})

	return nil
}
