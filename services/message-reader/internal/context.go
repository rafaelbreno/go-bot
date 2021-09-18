package internal

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Context struct {
	Logger *zap.Logger
	Env    map[string]string
}

func NewContext() *Context {
	l, _ := zap.NewProduction()

	if os.Getenv("APP_ENV") != "prod" {
		godotenv.Load()
	}

	return &Context{
		Logger: l,
		Env: map[string]string{
			"APP_URL":         os.Getenv("APP_URL"),
			"APP_PORT":        os.Getenv("APP_PORT"),
			"BOT_OAUTH_TOKEN": os.Getenv("BOT_OAUTH_TOKEN"),
			"BOT_NAME":        os.Getenv("BOT_NAME"),
			"REDIS_HOST":      os.Getenv("REDIS_HOST"),
			"REDIS_PORT":      os.Getenv("REDIS_PORT"),
		},
	}
}
