package main

import (
	"fmt"
	"log"
	"time"

	"github.com/vogo/duckdb"
	"gorm.io/gorm"
)

// Product represents a product entity with auto-incrementing ID
type Product struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Code      string    `gorm:"size:50;not null;uniqueIndex" json:"code"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Price     uint      `gorm:"not null" json:"price"`
	Category  string    `gorm:"size:50" json:"category"`
	InStock   bool      `gorm:"default:true" json:"in_stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User represents a user entity for demonstration
type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Age       int       `gorm:"check:age >= 0" json:"age"`
	Active    bool      `gorm:"default:true" json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	fmt.Println("🚀 Starting GORM DuckDB Example...")

	// 1. Establish database connection
	db, err := connectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("✅ Successfully connected to DuckDB")

	// 2. Auto-migrate tables
	if err := migrateDatabase(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("✅ Database migration completed")

	// 3. Perform CRUD operations
	if err := performCRUDOperations(db); err != nil {
		log.Fatal("CRUD operations failed:", err)
	}

	// 4. Demonstrate advanced queries
	if err := demonstrateAdvancedQueries(db); err != nil {
		log.Fatal("Advanced queries failed:", err)
	}

	// 5. Cleanup (GORM handles connection closing automatically)
	fmt.Println("\n🎉 All operations completed successfully!")
}

// connectDatabase establishes connection to DuckDB using GORM
func connectDatabase() (*gorm.DB, error) {
	// Connect to DuckDB (in-memory database)
	// For persistent database, use: duckdb.Open("example.ddb")
	db, err := gorm.Open(duckdb.Open(":memory:"), &gorm.Config{
		// Enable detailed logging for development
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DuckDB: %w", err)
	}

	return db, nil
}

// migrateDatabase creates tables with auto-incrementing primary keys
func migrateDatabase(db *gorm.DB) error {
	// Auto-migrate will create tables with proper schema
	err := db.AutoMigrate(&Product{}, &User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}

// performCRUDOperations demonstrates all CRUD operations
func performCRUDOperations(db *gorm.DB) error {
	fmt.Println("\n=== CRUD Operations Demo ===")

	// CREATE operations
	if err := createRecords(db); err != nil {
		return fmt.Errorf("create operations failed: %w", err)
	}

	// READ operations
	if err := readRecords(db); err != nil {
		return fmt.Errorf("read operations failed: %w", err)
	}

	// UPDATE operations
	if err := updateRecords(db); err != nil {
		return fmt.Errorf("update operations failed: %w", err)
	}

	// DELETE operations
	if err := deleteRecords(db); err != nil {
		return fmt.Errorf("delete operations failed: %w", err)
	}

	return nil
}

// createRecords demonstrates INSERT operations with auto-incrementing IDs
func createRecords(db *gorm.DB) error {
	fmt.Println("\n--- CREATE Operations ---")

	// Create single product
	product1 := Product{
		Code:     "P001",
		Name:     "Laptop",
		Price:    1200,
		Category: "Electronics",
		InStock:  true,
	}

	result := db.Create(&product1)
	if result.Error != nil {
		return fmt.Errorf("failed to create product: %w", result.Error)
	}
	fmt.Printf("✅ Created product with ID: %d\n", product1.ID)

	// Create multiple products in batch
	products := []Product{
		{Code: "P002", Name: "Mouse", Price: 25, Category: "Electronics", InStock: true},
		{Code: "P003", Name: "Keyboard", Price: 75, Category: "Electronics", InStock: true},
		{Code: "P004", Name: "Monitor", Price: 300, Category: "Electronics", InStock: false},
	}

	result = db.Create(&products)
	if result.Error != nil {
		return fmt.Errorf("failed to create products batch: %w", result.Error)
	}
	fmt.Printf("✅ Created %d products in batch\n", result.RowsAffected)

	// Create users
	users := []User{
		{Name: "Alice Johnson", Email: "alice@example.com", Age: 28, Active: true},
		{Name: "Bob Smith", Email: "bob@example.com", Age: 35, Active: true},
		{Name: "Charlie Brown", Email: "charlie@example.com", Age: 42, Active: false},
	}

	result = db.Create(&users)
	if result.Error != nil {
		return fmt.Errorf("failed to create users: %w", result.Error)
	}
	fmt.Printf("✅ Created %d users\n", result.RowsAffected)

	return nil
}

// readRecords demonstrates SELECT operations with various conditions
func readRecords(db *gorm.DB) error {
	fmt.Println("\n--- READ Operations ---")

	// Find all products
	var products []Product
	result := db.Find(&products)
	if result.Error != nil {
		return fmt.Errorf("failed to find products: %w", result.Error)
	}

	fmt.Printf("\n📦 All Products (%d found):\n", len(products))
	for _, p := range products {
		fmt.Printf("  ID: %d, Code: %s, Name: %s, Price: $%d, In Stock: %t\n",
			p.ID, p.Code, p.Name, p.Price, p.InStock)
	}

	// Find product by ID
	var product Product
	result = db.First(&product, 1)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("❌ Product with ID 1 not found")
		} else {
			return fmt.Errorf("failed to find product by ID: %w", result.Error)
		}
	} else {
		fmt.Printf("\n🔍 Product by ID 1: %s - $%d\n", product.Name, product.Price)
	}

	// Find products with conditions
	var expensiveProducts []Product
	result = db.Where("price > ?", 50).Find(&expensiveProducts)
	if result.Error != nil {
		return fmt.Errorf("failed to find expensive products: %w", result.Error)
	}
	fmt.Printf("\n💰 Expensive Products (>$50): %d found\n", len(expensiveProducts))

	// Find users with pagination
	var users []User
	result = db.Where("active = ?", true).Limit(10).Offset(0).Find(&users)
	if result.Error != nil {
		return fmt.Errorf("failed to find active users: %w", result.Error)
	}

	fmt.Printf("\n👥 Active Users (%d found):\n", len(users))
	for _, u := range users {
		fmt.Printf("  ID: %d, Name: %s, Email: %s, Age: %d\n",
			u.ID, u.Name, u.Email, u.Age)
	}

	return nil
}

// updateRecords demonstrates UPDATE operations
func updateRecords(db *gorm.DB) error {
	fmt.Println("\n--- UPDATE Operations ---")

	// Update single field
	result := db.Model(&Product{}).Where("code = ?", "P001").Update("price", 1100)
	if result.Error != nil {
		return fmt.Errorf("failed to update product price: %w", result.Error)
	}
	fmt.Printf("✅ Updated price for %d product(s)\n", result.RowsAffected)

	// Update multiple fields
	result = db.Model(&Product{}).Where("code = ?", "P002").Updates(Product{
		Name:  "Wireless Mouse",
		Price: 35,
	})
	if result.Error != nil {
		return fmt.Errorf("failed to update product: %w", result.Error)
	}
	fmt.Printf("✅ Updated %d product(s) with multiple fields\n", result.RowsAffected)

	// Update using map
	result = db.Model(&User{}).Where("email = ?", "alice@example.com").Updates(map[string]interface{}{
		"age":    29,
		"active": true,
	})
	if result.Error != nil {
		return fmt.Errorf("failed to update user: %w", result.Error)
	}
	fmt.Printf("✅ Updated %d user(s) using map\n", result.RowsAffected)

	// Batch update
	result = db.Model(&Product{}).Where("category = ?", "Electronics").Update("in_stock", true)
	if result.Error != nil {
		return fmt.Errorf("failed to batch update products: %w", result.Error)
	}
	fmt.Printf("✅ Batch updated %d product(s)\n", result.RowsAffected)

	return nil
}

// deleteRecords demonstrates DELETE operations
func deleteRecords(db *gorm.DB) error {
	fmt.Println("\n--- DELETE Operations ---")

	// Delete by primary key
	result := db.Delete(&Product{}, 4)
	if result.Error != nil {
		return fmt.Errorf("failed to delete product by ID: %w", result.Error)
	}
	fmt.Printf("✅ Deleted %d product(s) by ID\n", result.RowsAffected)

	// Delete with conditions
	result = db.Where("active = ?", false).Delete(&User{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete inactive users: %w", result.Error)
	}
	fmt.Printf("✅ Deleted %d inactive user(s)\n", result.RowsAffected)

	// Delete specific record
	var productToDelete Product
	db.Where("code = ?", "P003").First(&productToDelete)
	if productToDelete.ID != 0 {
		result = db.Delete(&productToDelete)
		if result.Error != nil {
			return fmt.Errorf("failed to delete specific product: %w", result.Error)
		}
		fmt.Printf("✅ Deleted product: %s\n", productToDelete.Name)
	}

	return nil
}

// demonstrateAdvancedQueries shows more complex database operations
func demonstrateAdvancedQueries(db *gorm.DB) error {
	fmt.Println("\n=== Advanced Queries Demo ===")

	// Count records
	var productCount int64
	db.Model(&Product{}).Count(&productCount)
	fmt.Printf("\n📊 Total products: %d\n", productCount)

	var userCount int64
	db.Model(&User{}).Where("active = ?", true).Count(&userCount)
	fmt.Printf("📊 Active users: %d\n", userCount)

	// Aggregate functions
	type Result struct {
		AvgPrice float64
		MaxPrice uint
		MinPrice uint
	}

	var result Result
	db.Model(&Product{}).Select("AVG(price) as avg_price, MAX(price) as max_price, MIN(price) as min_price").Scan(&result)
	fmt.Printf("\n💹 Price Statistics:\n")
	fmt.Printf("  Average: $%.2f\n", result.AvgPrice)
	fmt.Printf("  Maximum: $%d\n", result.MaxPrice)
	fmt.Printf("  Minimum: $%d\n", result.MinPrice)

	// Group by and having
	type CategoryStats struct {
		Category string
		Count    int64
		AvgPrice float64
	}

	var categoryStats []CategoryStats
	db.Model(&Product{}).Select("category, COUNT(*) as count, AVG(price) as avg_price").
		Group("category").Having("COUNT(*) > 0").Scan(&categoryStats)

	fmt.Printf("\n📈 Category Statistics:\n")
	for _, stat := range categoryStats {
		fmt.Printf("  %s: %d products, avg price $%.2f\n",
			stat.Category, stat.Count, stat.AvgPrice)
	}

	// Order and limit
	var topProducts []Product
	db.Where("in_stock = ?", true).Order("price DESC").Limit(3).Find(&topProducts)
	fmt.Printf("\n🏆 Top 3 Most Expensive In-Stock Products:\n")
	for i, p := range topProducts {
		fmt.Printf("  %d. %s - $%d\n", i+1, p.Name, p.Price)
	}

	return nil
}