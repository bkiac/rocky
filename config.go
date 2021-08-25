package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	key string
}

func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("KEY")

	return Config{
		key,
	}
}
