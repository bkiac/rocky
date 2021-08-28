package main

import (
	"os"

	"github.com/joho/godotenv"
)

var err = godotenv.Load()
var OMDBAPIKey = os.Getenv("OMDB_API_KEY")
var TelegramBotAPIKey = os.Getenv("TELEGRAM_BOT_API_KEY")
var Port = os.Getenv("PORT")
