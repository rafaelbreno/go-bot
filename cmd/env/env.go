package env

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go-bot/cmd/err"
)

type Env struct {
	Err *err.Err
}

// Loading .env file
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
		log.Fatalf("Message: \n %s", err.Error())
	}
}

func (e *Env) Getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		err := errors.New("Key not found")
		e.Err.Log(err)
	}
	return value
}
