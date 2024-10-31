package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/haiyen11231/eco-taxi-backend-user-service/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func ConnectToMySQL() error {
	// Building the Data Source Name (DSN) connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	// Opening a GORM connection with MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to MySQL:", err)
		return err
	}

	// Setting up connection pooling for sql.DB, which is underlying DB connection in GORM
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get sql.DB from GORM:", err)
		return err
	}

	// Configuring connection pool settings
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	DB = db
	DB.AutoMigrate(&model.User{})
	log.Println("Connected to MySQL!")
	
	return nil
}
