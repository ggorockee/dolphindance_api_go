package database

import (
	"fmt"
	"go_dev/config"
	"go_dev/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER", "root"),
		config.Config("DB_PASSWORD", "yourpassword"),
		config.Config("DB_HOST", "localhost"),
		config.Config("DB_PORT", "3306"),
		config.Config("DB_NAME", "your_database"),
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(
		&model.User{},
	)
	fmt.Println("Database Migrated")

}
