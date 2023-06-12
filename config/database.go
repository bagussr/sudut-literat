package config

import (
	"fmt"
	"sudut_literat/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable that holds the database connection
var DB *gorm.DB

// InitDB is a function that initializes the database connection
func InitDB() *gorm.DB {
	var err error
	dsn := "user=postgres password=dadasdudus12 dbname=sudut_literat port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	DB, err = gorm.Open( postgres.Open(dsn),&gorm.Config{})
	helper.Error(err)

	fmt.Println("Connection Opened to Database")
	return DB
}