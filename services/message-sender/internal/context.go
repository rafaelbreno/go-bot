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
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	l, _ := zap.NewProduction()

	c := Context{
		Env: map[string]string{
			"IRC_URL":         os.Getenv("API_IRC_URL"),
			"IRC_PORT":        os.Getenv("API_IRC_PORT"),
			"BOT_OAUTH_TOKEN": os.Getenv("BOT_OAUTH_TOKEN"),
			"BOT_NAME":        os.Getenv("BOT_NAME"),
		},
		Logger: l,
	}

	return &c
}
