package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//OpenDB connects to db
func OpenDB() *gorm.DB {
	if db != nil {
		if err := db.DB().Ping(); err != nil {
			db.DB().Close()
		} else {
			return db
		}
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSSL := os.Getenv("DB_SSL")

	connectionStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPass, dbSSL)
	driver := "postgres"

	db, err := gorm.Open(driver, connectionStr)
	if err != nil {
		log.Println("Failed to connect database "+connectionStr+". Error: %v", err)
		panic("OpenDB failed")
	}

	db.LogMode(true)
	log.Println("database connected!")
	return db
}
