package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rafaelbreno/go-bot/bot"
	"github.com/rafaelbreno/go-bot/conn"
	"github.com/rafaelbreno/go-bot/internal"
	"go.uber.org/zap"
)

var (
	connType   *string
	loadEnv    *string
	connection *conn.IRC
	ctx        *internal.Context
	logger     *zap.Logger
)

func newLogger() {
	logger, _ = zap.NewProduction()
}

func newConnection() *conn.IRC {
	c, err := conn.NewIRC(ctx)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(0)
	}

	return c
}

func init() {
	newLogger()

	connType = flag.String("conn", "IRC", "Connection type to Twitch's Chat (IRC ou WS)")
	loadEnv = flag.String("env", "", "If will load .env file")

	// Loading .env file
	if err := godotenv.Load(); err != nil {
		logger.Error(err.Error())
	}

	ctx = &internal.Context{
		Logger:     logger,
		Channels:   strings.Split(os.Getenv("CHANNEL_NAME"), ","),
		BotName:    os.Getenv("BOT_USERNAME"),
		OAuthToken: os.Getenv("BOT_OAUTH_TOKEN"),
	}
}

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	connection = newConnection()

	go bot.Start(ctx, connection)

	defer connection.Close()

	<-stop

	logger.Info("Gracefully terminating bot")
}
