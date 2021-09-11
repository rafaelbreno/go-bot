package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/storage"
	"go.uber.org/zap"
)

var (
	// Storage variable
	Storage *storage.Storage

	// Ctx - responsible for storing
	// Logger and env variables
	Ctx *internal.Context
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	l, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	Ctx = &internal.Context{
		Logger: l,
		Env: map[string]string{
			"PGSQL_HOST":     os.Getenv("PGSQL_HOST"),
			"PGSQL_PORT":     os.Getenv("PGSQL_PORT"),
			"PGSQL_USER":     os.Getenv("PGSQL_USER"),
			"PGSQL_PASSWORD": os.Getenv("PGSQL_PASSWORD"),
			"PGSQL_DBNAME":   os.Getenv("PGSQL_DBNAME"),
			"REDIS_HOST":     os.Getenv("REDIS_HOST"),
			"REDIS_PORT":     os.Getenv("REDIS_PORT"),
			"REDIS_PASSWORD": os.Getenv("REDIS_PASSWORD"),
			"KAFKA_URL":      os.Getenv("KAFKA_URL"),
			"KAFKA_PORT":     os.Getenv("KAFKA_PORT"),
			"KAFKA_TOPIC":    os.Getenv("KAFKA_TOPIC"),
			"AUTH_URL":       os.Getenv("AUTH_URL"),
			"AUTH_PORT":      os.Getenv("AUTH_PORT"),
			"API_PORT":       getPort(),
		},
	}

	Storage = storage.NewStorage(Ctx)
}

func getPort() string {
	envPort := os.Getenv("API_PORT")
	if []byte(envPort)[0] == ':' {
		return envPort
	}

	return string(append([]byte(":"), []byte(envPort)...))
}
