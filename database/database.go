package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const username = "sammarth"
const password = "mysql"
const databaseName = "gorm_test_DB"
const protocol = "@tcp(127.0.0.1:3306)/"

var DB *gorm.DB

func init() {

	dsn := username + ":" + password + protocol + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	handleError(err)

}

func handleError(err error) {
	fmt.Println(err)
}
