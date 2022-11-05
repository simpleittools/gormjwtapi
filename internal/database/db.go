package database

import (
	"fmt"
	"github.com/simpleittools/gormjwtapi/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

// This will define the DB

var DB *gorm.DB

func Conn() {
	dbType := os.Getenv("DB_ENGINE")
	dsn := ""
	switch dbType {
	case "POSTGRES":
		dsn = os.Getenv("POSTGRESDSN")
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Could not connect to the POSTGRES DB")
		} else {
			fmt.Println("connected to POSTGRES DB")
		}
		DB = conn
		conn.AutoMigrate(&models.User{})
	case "MYSQL":
		dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Could not connect to the MYSQL DB")
		} else {
			fmt.Println("connected to MYSQL")
		}
		DB = conn
		conn.AutoMigrate(&models.User{})
	case "SQLITE":
		conn, err := gorm.Open(sqlite.Open(os.Getenv("SQLITEDBNAME")), &gorm.Config{})
		if err != nil {
			log.Fatal("Could not connect to the SQLITE DB")
		} else {
			fmt.Println("connected to SQLITE")
		}
		DB = conn
		conn.AutoMigrate(&models.User{})
	default:
		panic("invalid database definition")
	}

}
