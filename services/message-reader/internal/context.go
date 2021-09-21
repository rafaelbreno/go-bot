package internal

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaelbreno/go-bot/services/message-reader/proto"
	"go.uber.org/zap"
)

type Context struct {
	Logger *zap.Logger
	Env    map[string]string
	Msg    *chan *proto.MessageRequest
}

func NewContext() *Context {
	l, _ := zap.NewProduction()

	if os.Getenv("APP_ENV") != "prod" {
		godotenv.Load()
	}

	msgChan := make(chan *proto.MessageRequest)

	return &Context{
		Logger: l,
		Env: map[string]string{
			"APP_URL":             os.Getenv("APP_URL"),
			"APP_PORT":            os.Getenv("APP_PORT"),
			"BOT_OAUTH_TOKEN":     os.Getenv("BOT_OAUTH_TOKEN"),
			"BOT_NAME":            os.Getenv("BOT_NAME"),
			"REDIS_HOST":          os.Getenv("REDIS_HOST"),
			"REDIS_PORT":          os.Getenv("REDIS_PORT"),
			"IRC_URL":             os.Getenv("API_IRC_URL"),
			"IRC_PORT":            os.Getenv("API_IRC_PORT"),
			"SENDER_SERVICE_URL":  os.Getenv("SENDER_SERVICE_URL"),
			"SENDER_SERVICE_PORT": os.Getenv("SENDER_SERVICE_PORT"),
		},
		Msg: &msgChan,
	}
}
