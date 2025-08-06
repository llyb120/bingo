package sql

import (
	"database/sql"

	"github.com/llyb120/bingo/core"
)

// 对外暴露的方法
var (
	Select func(db *sql.DB, sql string, args ...any) ([]map[string]string, error)
)

func _Select(db *sql.DB, sql string, args ...any) ([]map[string]string, error) {
	return nil, nil
}

func Init(state *core.State) {
	state.ExportFunc(&Select, _Select)
}
