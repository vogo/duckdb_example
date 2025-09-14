package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/marcboeker/go-duckdb/v2"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 连接到DuckDB数据库（内存模式）
	db, err := sql.Open("duckdb", "")
	if err != nil {
		log.Fatal("Failed to connect to DuckDB:", err)
	}
	defer db.Close()

	fmt.Println("✅ Successfully connected to DuckDB")

	// 创建表（包含自增ID）
	if err := createTable(db); err != nil {
		log.Fatal("Failed to create table:", err)
	}
	fmt.Println("✅ Table created successfully")

	// 插入数据
	if err := insertUsers(db); err != nil {
		log.Fatal("Failed to insert users:", err)
	}
	fmt.Println("✅ Users inserted successfully")

	// 查询所有用户
	if err := queryAllUsers(db); err != nil {
		log.Fatal("Failed to query users:", err)
	}

	// 查询特定用户
	if err := queryUserByID(db, 1); err != nil {
		log.Fatal("Failed to query user by ID:", err)
	}

	// 更新用户
	if err := updateUser(db, 1, "Alice Updated", 26); err != nil {
		log.Fatal("Failed to update user:", err)
	}
	fmt.Println("✅ User updated successfully")

	// 再次查询以验证更新
	if err := queryUserByID(db, 1); err != nil {
		log.Fatal("Failed to query updated user:", err)
	}

	// 删除用户
	if err := deleteUser(db, 2); err != nil {
		log.Fatal("Failed to delete user:", err)
	}
	fmt.Println("✅ User deleted successfully")

	// 最终查询所有用户
	fmt.Println("\n=== Final user list ===")
	if err := queryAllUsers(db); err != nil {
		log.Fatal("Failed to query final users:", err)
	}
}

// 创建用户表
func createTable(db *sql.DB) error {
	query := `
		CREATE SEQUENCE IF NOT EXISTS user_id_seq START 1;
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY DEFAULT nextval('user_id_seq'),
			name VARCHAR(100) NOT NULL,
			age INTEGER NOT NULL
		)
	`
	_, err := db.Exec(query)
	return err
}

// 插入用户数据
func insertUsers(db *sql.DB) error {
	users := []struct {
		name string
		age  int
	}{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	for _, user := range users {
		_, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.name, user.age)
		if err != nil {
			return fmt.Errorf("failed to insert user %s: %w", user.name, err)
		}
	}
	return nil
}

// 查询所有用户
func queryAllUsers(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, age FROM users ORDER BY id")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("\n=== All Users ===")
	fmt.Printf("%-5s %-15s %-5s\n", "ID", "Name", "Age")
	fmt.Println("------------------------------")

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return err
		}
		fmt.Printf("%-5d %-15s %-5d\n", user.ID, user.Name, user.Age)
	}

	return rows.Err()
}

// 根据ID查询用户
func queryUserByID(db *sql.DB, id int) error {
	var user User
	err := db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("\n❌ No user found with ID: %d\n", id)
			return nil
		}
		return err
	}

	fmt.Printf("\n=== User ID %d ===\n", id)
	fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	return nil
}

// 更新用户信息
func updateUser(db *sql.DB, id int, name string, age int) error {
	result, err := db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?", name, age, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with ID: %d", id)
	}

	return nil
}

// 删除用户
func deleteUser(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with ID: %d", id)
	}

	return nil
}
