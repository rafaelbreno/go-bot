package env

import (
	"fmt"
	"log"
	"os"

	"go-bot/cmd/app_error"

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

func (e *Env) Getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		errMsg := fmt.Sprintf(`Key "%s" not found!`, key)
		app_error.NewError(errMsg, "cmd/env.go:29")
	}
	return value
}
