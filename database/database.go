package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

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

	// When running without docker => cfg.host = "localhost"
	cfg.host = "host.docker.internal"

	cfg.protocol = "@tcp(" + cfg.host + ":" + cfg.port + ")/"

	return cfg
}

func handleError(err error) {
	fmt.Println(err)
}
