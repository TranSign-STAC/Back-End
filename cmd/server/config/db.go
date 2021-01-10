package config

import (
	"fmt"

	"transign/cmd/server/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB has the DB interface
var DB *gorm.DB

// ConnectDB connects to the postgres database
func ConnectDB() *gorm.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", Envs["DB_HOST"], Envs["DB_PORT"], Envs["DB_USERNAME"], Envs["DB_PASSWORD"], Envs["DB_DATABASE"], Envs["DB_SSLMODE"])
	var err error
	DB, err = gorm.Open(postgres.Open(dbInfo), &gorm.Config{})

	if err != nil {
		logrus.Fatal(err)
		panic("Failed to connect to the database.")
	}

	DB.AutoMigrate(&models.Translation{})
	return DB
}
