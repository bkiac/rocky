package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	omdbAPIKey        string
	TelegramBotAPIKey string
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	omdbAPIKey := os.Getenv("OMDB_API_KEY")
	telegramBotAPIKey := os.Getenv("TELEGRAM_BOT_API_KEY")

	return Config{
		omdbAPIKey,
		telegramBotAPIKey,
	}
}
