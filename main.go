package main

import (
	"log"

	"github.com/joho/godotenv"
)

// func getKey(key string) string {
// 	env, _ := os.Open(".env")

// 	return env.
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
