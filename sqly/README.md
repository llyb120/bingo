# Sqly - 数据库查询模块

Sqly 是 Bingo 框架的数据库查询模块，提供了简洁的泛型化 SQL 查询功能。

## 功能特性

- **泛型支持**: 使用 `y.Cast` 自动将查询结果转换为指定类型 T
- **统一接口**: 只有一个 `Select` 函数，支持所有查询需求
- **原始数据**: 先转换为 `[][]any` 格式，再使用 `y.Cast` 进行类型转换
- **事务支持**: 支持数据库事务操作
- **错误处理**: 完善的错误处理和日志记录

## 核心原理

`sqly.Select` 的工作流程：
1. **执行SQL查询** → 获取数据库行和列信息
2. **驱动类型转换** → 将数据库驱动特有类型转换为Go基本类型
3. **标准化为 `[]map[string]any`** → 保持字段名的统一数据格式  
4. **使用 `y.Cast`** → 转换为目标类型 T

`sqly.Exec` 的工作流程：
1. **执行SQL语句** → INSERT、UPDATE、DELETE等操作
2. **返回完整结果** → 受影响的行数、自增ID（如果有）、错误信息

### 类型转换详解

数据库扫描出的数据通常不是Go基本类型，而是数据库驱动特有的类型：
- `[]byte` → 根据数据库列类型转换为 `string`、`int64`、`float64`、`bool`、`time.Time` 等
- `sql.NullString`、`sql.NullInt64` 等 → 提取有效值或返回 `nil`
- `driver.Valuer` → 调用 `Value()` 方法获取底层值
- 其他未知类型 → 使用反射获取底层接口值

将数据转换为 `[]map[string]any` 格式的优势：
- **保持字段名** → 列名信息不会丢失，便于结构体映射
- **类型安全** → 确保都是Go基本类型，提高 `y.Cast` 转换成功率
- **灵活性** → 支持动态字段和复杂的数据结构转换

## 使用示例

### 查询单个记录

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// 查询单个用户
user, err := sqly.Select[User](ctx, db, "SELECT id, name, age FROM users WHERE id = ?", 1)
if err != nil {
    log.Error(ctx, "查询失败: %v", err)
    return
}
```

### 查询多条记录

```go
// 查询用户列表
users, err := sqly.Select[[]User](ctx, db, "SELECT id, name, age FROM users WHERE age > ?", 18)
if err != nil {
    log.Error(ctx, "查询失败: %v", err)
    return
}
```

### 查询基本类型

```go
// 查询数量
count, err := sqly.Select[int](ctx, db, "SELECT COUNT(*) FROM users")

// 查询字符串
name, err := sqly.Select[string](ctx, db, "SELECT name FROM users WHERE id = ?", 1)
```

### 查询Map类型

```go
// 查询单个记录为Map
userMap, err := sqly.Select[map[string]interface{}](ctx, db, "SELECT * FROM users WHERE id = ?", 1)

// 查询多条记录为Map数组
userMaps, err := sqly.Select[[]map[string]interface{}](ctx, db, "SELECT * FROM users")
```

### 查询原始数据

```go
// 获取原始 []map[string]any 数据（保持字段名）
rawData, err := sqly.Select[[]map[string]any](ctx, db, "SELECT * FROM users")
if err != nil {
    return err
}

// 可以进一步处理 rawData
for _, row := range rawData {
    for column, value := range row {
        fmt.Printf("%s: %v ", column, value)
    }
    fmt.Println()
}
```

### 执行SQL语句

`sqly.Exec` 返回三个值：`(受影响行数, 自增ID, 错误)`

#### 插入操作（获取自增ID）

```go
// 插入数据并获取自增ID
rowsAffected, lastInsertId, err := sqly.Exec(ctx, db, "INSERT INTO users (name, age) VALUES (?, ?)", "张三", 25)
if err != nil {
    return err
}
fmt.Printf("插入了 %d 条记录，新记录ID: %d\n", rowsAffected, lastInsertId)

// 批量插入（多次使用lastInsertId）
for i, user := range users {
    rowsAffected, lastInsertId, err := sqly.Exec(ctx, db, "INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
    if err != nil {
        return fmt.Errorf("插入第%d个用户失败: %w", i+1, err)
    }
    fmt.Printf("用户 %s 插入成功，ID: %d\n", user.Name, lastInsertId)
}
```

#### 更新和删除操作（使用下划线忽略自增ID）

```go
// 更新数据（不需要自增ID）
rowsAffected, _, err := sqly.Exec(ctx, db, "UPDATE users SET age = ? WHERE name = ?", 26, "张三")
if err != nil {
    return err
}
if rowsAffected == 0 {
    fmt.Println("没有找到匹配的记录进行更新")
} else {
    fmt.Printf("更新了 %d 条记录\n", rowsAffected)
}

// 删除数据（不需要自增ID）
rowsAffected, _, err = sqly.Exec(ctx, db, "DELETE FROM users WHERE age > ?", 60)
if err != nil {
    return err
}
fmt.Printf("删除了 %d 条记录\n", rowsAffected)
```

#### 事务中的操作

```go
tx, err := sqly.NewTx(db)
if err != nil {
    return err
}
defer tx.Rollback()

// 在事务中插入并获取ID
_, userId, err := sqly.Exec(ctx, tx, "INSERT INTO users (name, age) VALUES (?, ?)", "李四", 30)
if err != nil {
    return err
}

// 使用获取到的ID插入关联数据
_, _, err = sqly.Exec(ctx, tx, "INSERT INTO user_profiles (user_id, bio) VALUES (?, ?)", userId, "这是李四的简介")
if err != nil {
    return err
}

// 提交事务
return tx.Commit()
```

### 事务使用

```go
// 创建事务
tx, err := sqly.NewTx(db)
if err != nil {
    return err
}
defer tx.Rollback()

// 在事务中查询
user, err := sqly.Select[User](ctx, tx, "SELECT * FROM users WHERE id = ?", 1)
if err != nil {
    return err
}

// 提交事务
return tx.Commit()
```

### 与 Gox 结合使用

```go
// 使用 gox 生成的查询
query := sql.TestGox() // 返回 gox.Query

// 执行查询并转换成目标类型
users, err := sqly.Select[[]User](ctx, db, query.SQL, query.Args...)
if err != nil {
    log.Error(ctx, "查询失败: %v", err)
    return
}
```

## 结构体映射

使用 `json` 标签进行字段映射（配合 `y.Cast` 的转换规则）：

```go
type User struct {
    ID       int    `json:"id"`         // 映射数据库的 id 列
    Name     string `json:"name"`       // 映射数据库的 name 列  
    FullName string `json:"full_name"`  // 映射数据库的 full_name 列
    Age      int    `json:"age"`
}
```

## 支持的类型转换

借助 `y.Cast` 的强大类型转换能力，支持：

1. **基本类型**: `int`, `string`, `float64`, `bool` 等
2. **结构体**: 根据 `json` 标签自动映射
3. **数组/切片**: `[]User`, `[]map[string]interface{}`, `[]map[string]any` 等
4. **Map**: `map[string]interface{}`, `map[string]any` 等
5. **嵌套结构**: 复杂的嵌套数据结构

## Exec方法返回值说明

| 返回值 | 类型 | 说明 |
|--------|------|------|
| 受影响行数 | `int64` | INSERT、UPDATE、DELETE操作影响的记录数量 |
| 自增ID | `int64` | INSERT操作时的自增主键值，其他操作通常为0 |
| 错误 | `error` | 操作失败时的错误信息，成功时为nil |

**注意事项：**
- 自增ID只在INSERT操作且表有自增主键时有效
- 某些数据库驱动可能不支持LastInsertId，此时返回0
- 批量插入时，LastInsertId通常返回第一条记录的ID

## 错误处理

- 数据库连接错误
- SQL 语法错误  
- 类型转换错误
- 查询执行错误

所有错误都会通过日志记录并返回详细的错误信息。