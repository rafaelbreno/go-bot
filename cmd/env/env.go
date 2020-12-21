package env

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct{}

// Loading .env file
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
		log.Fatalf("Message: \n %s", err.Error())
	}
}

func (_ *Env) Getenv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("Key not found")
	}
	return value, nil
}
