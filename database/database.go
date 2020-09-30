package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type URLData struct {
	ID               string `gorm:"primaryKey"`
	URL              string
	CrawlTimeout     int
	Frequency        int
	FailureThreshold int
	Status           string
	FailureCount     int
}

type envConfig struct {
	host, port, username, password, databaseName, protocol string
}

// To access the database
var DB *gorm.DB

func InitializeDatabase() {

	cfg := getConfig()
	dsn := cfg.username + ":" + cfg.password + cfg.protocol + cfg.databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(cfg)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	handleError(err)

	err = DB.AutoMigrate(&URLData{}) // Makes the table of structure URLData
	if err != nil {
		os.Exit(1)
	}
}

// Builds the config to setup the database
func getConfig() envConfig {

	err := godotenv.Load(".env")
	handleError(err)

	cfg := envConfig{
		port:         "3306",
		username:     os.Getenv("DB_USER"),
		password:     os.Getenv("DB_PASSWORD"),
		databaseName: os.Getenv("DB_NAME"),
	}

	isDocker, _ := os.LookupEnv("DOCKER")

	if isDocker == "false" {
		cfg.host = "localhost"
	} else {
		cfg.host = "host.docker.internal"
	}

	cfg.protocol = "@tcp(" + cfg.host + ":" + cfg.port + ")/"

	return cfg
}

func handleError(err error) {
	fmt.Println(err)
}
