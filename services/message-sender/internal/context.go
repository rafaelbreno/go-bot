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

	c := Context{
		Env: map[string]string{
			"IRC_URL":  os.Getenv("API_IRC_URL"),
			"IRC_PORT": os.Getenv("API_IRC_PORT"),
		},
		Logger: zap.NewProduction(),
	}

	return &c
}
