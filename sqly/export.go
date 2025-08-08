package sqly

import (
	"fmt"

	"github.com/llyb120/bingo/core"
)

// 对外暴露的方法
var (
	Starter core.Starter = &sqlyStarter{}

	Select func(db DBLike, sql string, args ...any) ([]map[string]string, error)
)

func _Select(db DBLike, sql string, args ...any) ([]map[string]string, error) {
	fmt.Println("滤哥威武")
	return nil, nil
}

type sqlyStarter struct {
}

func (s *sqlyStarter) Init(state *core.State) {
	core.ExportFunc(state, "sql.Select", &Select, _Select)
}

func (s *sqlyStarter) Destroy(state *core.State) {

}
