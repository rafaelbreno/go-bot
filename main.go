package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot_username := os.Getenv("BOT_USERNAME")
	channel_name := os.Getenv("CHANNEL_NAME")
	oauth_key := os.Getenv("OAUTH_TOKEN")

}
