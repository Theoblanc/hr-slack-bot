package main

import (
	"fmt"
	"log"

	"github.com/TheoBlanc/AttendEase/config"
	"github.com/TheoBlanc/AttendEase/internal/infra"
	"github.com/TheoBlanc/AttendEase/internal/infra/database"
)

func main() {
	config := config.Load()

	_, err := database.PostgresDBInitialize(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	slackClient := infra.SlackInitialize(config.SlackBotToken)

	fmt.Println("Welcome to AttendEase - Your attendance and leave management system!")

}
