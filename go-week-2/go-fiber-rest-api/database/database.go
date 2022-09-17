package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	const mysqldb = "go-fiber_gorm:golang-gofiber@tcp(127.0.0.1:3306)/gofiber_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := mysqldb
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Successfully connect to database")
}
