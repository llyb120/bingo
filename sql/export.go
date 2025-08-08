package sql

import (
	"github.com/llyb120/bingo/core"
)

// 对外暴露的方法
var (
	Select func(db DBLike, sql string, args ...any) ([]map[string]string, error)
)

func _Select(db DBLike, sql string, args ...any) ([]map[string]string, error) {
	return nil, nil
}

func Init(state *core.State) {
	core.ExportFunc(state, &Select, _Select)
}
