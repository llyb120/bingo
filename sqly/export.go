package sqly

import (
	"fmt"
	"time"

	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/log"
)

//var (
//	Select = core.ExportFunc("sqly.Select", _Select)
//)

func Select(ctx core.Context, db DBLike, sql string, args ...any) ([]map[string]string, error) {
	log.Info(ctx, "滤哥威武")
	return nil, fmt.Errorf("oh shit")
}

func TestSelect(ctx core.Context, ok string) (string, error) {
	time.Sleep(2 * time.Second)
	return ok + "fu", nil
}
