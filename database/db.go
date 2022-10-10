package database

import (
	"assignment_02/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "khairulazman"
	password = "6720"
	dbPort   = "5432"
	dbname   = "assignment_02"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user,
		password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database :", err)
	}
	fmt.Println("Successfully connected to database")

	db.Debug().AutoMigrate(models.Items{}, models.Orders{})
}

func GetDB() *gorm.DB {
	return db
}
