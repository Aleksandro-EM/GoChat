package main

import (
	"GoChat/database"
	"GoChat/models"
	"log"
)

func main() {
	db := database.Connect()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Unable to migrate users table")
	}
	print("Connected to Database!")
}
