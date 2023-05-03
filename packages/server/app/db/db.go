package db

import (
	"log"
	"os"

	"go.labs/server/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL, dbURLExists := os.LookupEnv("dbURL")

	if !dbURLExists {
		dbURL = "postgres://postgres:postgres@localhost:5432/discusser"
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Account{}, &models.Token{}, &models.Post{}, &models.Message{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
