package config

import (
	"fmt"
	"os"
)

type Config struct {
	SlackClientID  string
	SlackChannelID string
	SlackBotToken  string
	DBHost         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBPort         string
	DatabaseURL    string
}

func Load() *Config {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	return &Config{
		DatabaseURL:    databaseURL,
		SlackClientID:  os.Getenv("SLACK_CLIENT_ID"),
		SlackChannelID: os.Getenv("SLACK_CHANNEL_ID"),
		SlackBotToken:  os.Getenv("SLACK_BOT_TOKEN"),
	}
}
