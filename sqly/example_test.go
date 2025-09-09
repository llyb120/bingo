package sqly_test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/llyb120/bingo/core"
	"github.com/llyb120/bingo/sqly"
)

// 示例用户结构体
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// TestSelectExamples 展示如何使用 sqly.Select 功能
func TestSelectExamples(t *testing.T) {
	// 注意：这些是示例代码，需要实际的数据库连接才能运行
	// 在实际使用中，您需要配置数据库连接

	var db *sql.DB
	var ctx core.Context

	// 示例1: 查询单个用户（结构体）
	// SQL查询 -> [][]any -> y.Cast -> User
	user, err := sqly.Select[User](ctx, db, "SELECT id, name, email, age FROM users WHERE id = ?", 1)
	if err != nil {
		t.Logf("查询用户失败: %v", err)
	}
	_ = user

	// 示例2: 查询用户列表（结构体数组）
	// SQL查询 -> [][]any -> y.Cast -> []User
	users, err := sqly.Select[[]User](ctx, db, "SELECT id, name, email, age FROM users WHERE age > ?", 18)
	if err != nil {
		t.Logf("查询用户列表失败: %v", err)
	}
	_ = users

	// 示例3: 查询用户数量（基本类型）
	// SQL查询 -> [][]any -> y.Cast -> int
	count, err := sqly.Select[int](ctx, db, "SELECT COUNT(*) FROM users")
	if err != nil {
		t.Logf("查询用户数量失败: %v", err)
	}
	_ = count

	// 示例4: 查询用户信息（Map）
	// SQL查询 -> [][]any -> y.Cast -> map[string]interface{}
	userMap, err := sqly.Select[map[string]interface{}](ctx, db, "SELECT * FROM users WHERE id = ?", 1)
	if err != nil {
		t.Logf("查询用户信息失败: %v", err)
	}
	_ = userMap

	// 示例5: 查询原始数据（[]map[string]any）
	// SQL查询 -> []map[string]any（保持字段名）
	rawData, err := sqly.Select[[]map[string]any](ctx, db, "SELECT * FROM users")
	if err != nil {
		t.Logf("查询原始数据失败: %v", err)
	}
	_ = rawData

	// 示例6: 执行SQL语句
	// 插入数据
	rowsAffected, lastInsertId, err := sqly.Exec(ctx, db, "INSERT INTO users (name, email, age) VALUES (?, ?, ?)", "张三", "zhangsan@test.com", 25)
	if err != nil {
		t.Logf("插入数据失败: %v", err)
	}
	t.Logf("插入了 %d 条记录，新记录ID: %d", rowsAffected, lastInsertId)

	// 更新数据
	rowsAffected, _, err = sqly.Exec(ctx, db, "UPDATE users SET age = ? WHERE name = ?", 26, "张三")
	if err != nil {
		t.Logf("更新数据失败: %v", err)
	}
	t.Logf("更新了 %d 条记录", rowsAffected)

	// 删除数据
	rowsAffected, _, err = sqly.Exec(ctx, db, "DELETE FROM users WHERE age > ?", 60)
	if err != nil {
		t.Logf("删除数据失败: %v", err)
	}
	t.Logf("删除了 %d 条记录", rowsAffected)

	// 示例7: 使用事务
	tx, err := sqly.NewTx(db)
	if err != nil {
		t.Logf("创建事务失败: %v", err)
		return
	}
	defer tx.Rollback()

	// 在事务中查询
	txUser, err := sqly.Select[User](ctx, tx, "SELECT * FROM users WHERE id = ?", 1)
	if err != nil {
		t.Logf("事务查询失败: %v", err)
	}
	_ = txUser

	// 提交事务
	if err := tx.Commit(); err != nil {
		t.Logf("提交事务失败: %v", err)
	}
}

// TestSelectWithGox 展示如何与 gox 结合使用
func TestSelectWithGox(t *testing.T) {
	// 这个测试展示了如何将 gox 生成的查询与 sqly 结合使用
	// 在实际项目中，您可以这样使用：

	/*
		// 假设您有一个 gox 查询函数
		query := sql.TestGox() // 返回 gox.Query

		// 执行查询并转换成目标类型
		users, err := sqly.Select[[]User](ctx, db, query.SQL, query.Args...)
		if err != nil {
			log.Error(ctx, "查询失败: %v", err)
			return
		}
	*/
}
