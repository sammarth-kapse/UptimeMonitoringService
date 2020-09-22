package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const username = "sammarth"
const password = "mysql"
const databaseName = "gorm_test_db"
const protocol = "@tcp(127.0.0.1:3306)/"
const INACTIVE = "inactive"
const ACTIVE = "active"

type URLData struct {
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

	dsn := username + ":" + password + protocol + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	handleError(err)

	err = db.AutoMigrate(&URLData{})
	handleError(err)

}

func (urlInfo *URLData) getURLInfoFromDatabase() {
	db.First(&urlInfo)
}

func (urlInfo *URLData) saveIntoDatabase() {
	db.Save(&urlInfo)
}

// Utility Functions:

func checkIfURLEmpty(urlInfo URLData) bool {
	return urlInfo.URL == ""
}

func handleError(err error) {
	fmt.Println(err)
}
