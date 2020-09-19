package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const username = "sammarth"
const password = "mysql"
const INACTIVE = "inactive"
const ACTIVE = "active"

type UrlData struct {
	ID               string `gorm:"primaryKey"`
	URL              string
	CrawlTimeout     int
	Frequency        int
	FailureThreshold int
	Status           string
	FailureCount     int
}

var db *gorm.DB

func init() {

	dsn := username + ":" + password + "@tcp(127.0.0.1:3306)/gorm_test_db?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	handleError(err)

	err = db.AutoMigrate(&UrlData{})
	handleError(err)

}

// Utility Functions:

func checkIfURLEmpty(urlInfo UrlData) bool {
	return urlInfo.URL == ""
}

func handleError(err error) {
	fmt.Println(err)
}
