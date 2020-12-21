package env

import (
	"errors"
	"fmt"
	"log"
	"os"

	"go-bot/cmd/err"

	"github.com/joho/godotenv"
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
		errMsg := fmt.Sprintf(`Key "%s" not found!`, key)
		err := errors.New(errMsg)
		e.Err.Log(err)
	}
	return value
}
