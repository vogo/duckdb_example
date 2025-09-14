# DuckDB Go Example

这是一个使用Go语言和DuckDB数据库的完整示例，演示了基本的CRUD操作（创建、读取、更新、删除）以及自增ID功能。

## 功能特性

- ✅ 数据库连接（内存模式）
- ✅ 创建表结构（带自增ID）
- ✅ 插入数据
- ✅ 查询所有记录
- ✅ 根据ID查询单条记录
- ✅ 更新记录
- ✅ 删除记录

## 依赖

- Go 1.25.0+
- [go-duckdb](https://github.com/marcboeker/go-duckdb) v2.3.3

## 安装和运行

1. 进入项目目录：
```bash
cd official
```

2. 下载依赖：
```bash
go mod tidy
```

3. 运行示例：
```bash
go run main.go
```

## 示例输出

```
✅ Successfully connected to DuckDB
✅ Table created successfully
✅ Users inserted successfully

=== All Users ===
ID    Name            Age  
------------------------------
1     Alice           25   
2     Bob             30   
3     Charlie         35   

=== User ID 1 ===
ID: 1, Name: Alice, Age: 25
✅ User updated successfully

=== User ID 1 ===
ID: 1, Name: Alice Updated, Age: 26
✅ User deleted successfully

=== Final user list ===

=== All Users ===
ID    Name            Age  
------------------------------
1     Alice Updated   26   
3     Charlie         35   
```

## 代码结构

### 主要函数

- `createTable()` - 创建用户表，包含自增ID序列
- `insertUsers()` - 批量插入用户数据
- `queryAllUsers()` - 查询并显示所有用户
- `queryUserByID()` - 根据ID查询特定用户
- `updateUser()` - 更新用户信息
- `deleteUser()` - 删除用户

### 数据结构

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

### 表结构

```sql
CREATE SEQUENCE IF NOT EXISTS user_id_seq START 1;
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY DEFAULT nextval('user_id_seq'),
    name VARCHAR(100) NOT NULL,
    age INTEGER NOT NULL
)
```

## 技术要点

1. **自增ID实现**：使用DuckDB的SEQUENCE和DEFAULT nextval()来实现自增主键
2. **错误处理**：完善的错误处理机制，包括数据库连接、SQL执行等
3. **资源管理**：正确使用defer关闭数据库连接和查询结果
4. **参数化查询**：使用占位符防止SQL注入
5. **事务安全**：每个操作都有适当的错误检查和回滚机制