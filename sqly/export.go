package sqly

import (
	"fmt"

	"github.com/llyb120/bingo/core"
)

// 对外暴露的方法
var (
	Select = core.ExportFunc("sql.Select", _Select)
)

func _Select(db DBLike, sql string, args ...any) ([]map[string]string, error) {
	fmt.Println("滤哥威武")
	return nil, nil
}
