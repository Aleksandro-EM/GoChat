package main

import (
	"GoChat/database"
	"GoChat/models"
)

func main() {
	db := database.Connect()

	// Auto-migrate all tables
	db.AutoMigrate(
		&models.User{},
		&models.Chat{},
		&models.ChatMember{},
		&models.Message{},
		&models.Webhook{},
		&models.WebhookLog{},
	)
	print("Connected to Database!")
}
