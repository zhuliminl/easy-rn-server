package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zhuliminl/easyrn-server/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Init() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	log.Println("SetupDatabaseConnectionWithDSN", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	// 数据迁移
	db.AutoMigrate(&entity.User{}, &entity.Book{})
	DB = db
}

func CloseDatabaseConnection() {
	dbSQL, err := DB.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}