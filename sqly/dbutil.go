package sqly

import (
	"fmt"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/log"
	"github.com/llyb120/yoya2/y"
)

type DBUtil interface {
	Exec(ctx core.Context, db DBLike, sql string, args ...any) (int64, int64, error)
	Select(ctx core.Context, db DBLike, result any, sql string, args ...any) error
}

var (
	_ DBUtil = (*defaultDBUtil)(nil)
)

type defaultDBUtil struct {
}

func (d *defaultDBUtil) Exec(ctx core.Context, db DBLike, sql string, args ...any) (int64, int64, error) {
	//log.Info(ctx, "执行SQL语句: %s", sql)
	if db == nil {
		return 0, 0, fmt.Errorf("数据库连接不能为空")
	}

	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Error(ctx, "SQL执行失败: %v", err)
		return 0, 0, fmt.Errorf("SQL执行失败: %w", err)
	}

	// 获取受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error(ctx, "获取影响行数失败: %v", err)
		return 0, 0, fmt.Errorf("获取影响行数失败: %w", err)
	}

	// 获取自增ID（如果有）
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		// 某些数据库驱动或操作不支持LastInsertId，这是正常的
		log.Debug(ctx, "获取自增ID失败（可能不支持或不适用）: %v", err)
		lastInsertId = 0 // 设置为0表示没有自增ID或不支持
	}

	//log.Info(ctx, "SQL执行成功，影响行数: %d，自增ID: %d", rowsAffected, lastInsertId)
	return rowsAffected, lastInsertId, nil
}

func (d *defaultDBUtil) Select(ctx core.Context, db DBLike, result any, sql string, args ...any) error {
	//var zero T

	//log.Info(ctx, "执行SQL查询: \n%s%s%s \n", log.ColorYellow, PrintSql(ctx, sql, args...), log.ColorReset)

	if db == nil {
		return fmt.Errorf("数据库连接不能为空")
	}

	// 执行查询获取map格式数据（保持字段名）
	data, err := queryToMaps(ctx, db, sql, args...)
	if err != nil {
		return err
	}

	// 使用y.Cast转换成目标类型
	//var result T
	err = y.Cast(result, data)
	if err != nil {
		log.Error(ctx, "类型转换失败: %v", err)
		return fmt.Errorf("类型转换失败: %w", err)
	}

	return nil
}
