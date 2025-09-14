# GORM DuckDB Driver Example

这是一个使用 [github.com/vogo/duckdb](https://github.com/vogo/duckdb) GORM 驱动的综合示例，展示了如何在 Go 应用中使用 DuckDB 进行数据库操作。

## 功能特性

- ✅ **数据库连接**: 使用 GORM 连接到 DuckDB
- ✅ **自动迁移**: 自动创建表结构
- ✅ **自增ID**: 支持主键自增
- ✅ **CRUD操作**: 完整的增删改查功能
- ✅ **批量操作**: 批量插入和更新
- ✅ **高级查询**: 聚合查询、统计分析
- ✅ **错误处理**: 完善的错误处理机制
- ✅ **资源清理**: 自动关闭数据库连接

## 依赖项

- Go 1.25.0+
- [gorm.io/gorm](https://gorm.io/) v1.25.12
- [github.com/vogo/duckdb](https://github.com/vogo/duckdb) (GORM DuckDB 驱动)

## 安装和运行

1. **克隆或创建项目目录**:
   ```bash
   mkdir gorm_driver && cd gorm_driver
   ```

2. **初始化 Go 模块**:
   ```bash
   go mod init github.com/vogo/duckdb_example/gorm_driver
   ```

3. **安装依赖**:
   ```bash
   go mod tidy
   ```

4. **运行示例**:
   ```bash
   go run main.go
   ```

## 示例输出

```
🚀 Starting GORM DuckDB Example...
✅ Successfully connected to DuckDB
✅ Database migration completed

=== CRUD Operations Demo ===

--- CREATE Operations ---
✅ Created product with ID: 1
✅ Created 3 products in batch
✅ Created 3 users

--- READ Operations ---
📦 All Products (4 found):
  ID: 1, Code: P001, Name: Laptop, Price: $1200, In Stock: true
  ...

--- UPDATE Operations ---
✅ Updated price for 1 product(s)
✅ Updated 1 product(s) with multiple fields
...

🎉 All operations completed successfully!
```

## 代码结构

### 主要组件

- **数据模型**: `Product` 和 `User` 结构体
- **数据库连接**: 使用 GORM 连接 DuckDB
- **CRUD操作**: 完整的数据库操作函数
- **高级查询**: 聚合和统计查询

### 核心函数

- `connectDB()`: 建立数据库连接
- `createOperations()`: 创建数据记录
- `readOperations()`: 查询数据记录
- `updateOperations()`: 更新数据记录
- `deleteOperations()`: 删除数据记录
- `advancedQueries()`: 高级查询和统计

### 数据模型

```go
type Product struct {
    ID       uint   `gorm:"primaryKey;autoIncrement"`
    Code     string `gorm:"uniqueIndex;size:50;not null"`
    Name     string `gorm:"size:100;not null"`
    Price    float64
    InStock  bool `gorm:"default:true"`
    Category string `gorm:"size:50;default:'Electronics'"`
}

type User struct {
    ID     uint   `gorm:"primaryKey;autoIncrement"`
    Name   string `gorm:"size:100;not null"`
    Email  string `gorm:"uniqueIndex;size:100;not null"`
    Age    int
    Active bool `gorm:"default:true"`
}
```

## 技术要点

### GORM 特性
- **自动迁移**: 使用 `AutoMigrate()` 自动创建表
- **模型标签**: 使用 GORM 标签定义字段约束
- **关联查询**: 支持复杂的数据库查询
- **事务支持**: 内置事务处理机制

### DuckDB 优势
- **内存数据库**: 高性能的列式存储
- **SQL 兼容**: 支持标准 SQL 语法
- **分析友好**: 优化的 OLAP 查询性能
- **零配置**: 无需额外的数据库服务器

### 错误处理
- 所有数据库操作都包含错误检查
- 使用 `panic` 处理致命错误
- 详细的错误日志输出

## 扩展建议

1. **添加更多模型**: 创建更复杂的数据关系
2. **实现 API**: 将数据库操作封装为 REST API
3. **添加测试**: 编写单元测试和集成测试
4. **性能优化**: 使用索引和查询优化
5. **配置管理**: 使用配置文件管理数据库连接

## 相关资源

- [GORM 官方文档](https://gorm.io/docs/)
- [DuckDB 官方网站](https://duckdb.org/)
- [github.com/vogo/duckdb](https://github.com/vogo/duckdb)
- [Go 官方文档](https://golang.org/doc/)