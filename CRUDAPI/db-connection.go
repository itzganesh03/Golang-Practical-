package main

import (
	"fmt"

	// "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:admin@tcp(localhost:3306)/golang_assignment?parseTime=true"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("connection failde")
	}
	Database.AutoMigrate(&User{})
}
