package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define the User struct
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"uniqueIndex;size:255"`
	Password string
}

func main() {
	// Database connection string
	dsn := "host=localhost port=5432 dbname=pouyan user=postgres password=43754375 connect_timeout=10 sslmode=prefer TimeZone=Asia/Shanghai"
	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	// Automatically migrate your schema
	db.AutoMigrate(&User{})

	// Create a new user
	user := User{Name: "John Doe", Email: "john@example.com", Password: "password123"}
	db.Create(&user)

	// Query the user
	var queriedUser User
	db.First(&queriedUser, "email = ?", "john@example.com")
	log.Println("Queried User:", queriedUser)
}
