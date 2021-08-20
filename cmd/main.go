package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rafaelbreno/go-bot/bot"
	"github.com/rafaelbreno/go-bot/conn"
	"github.com/rafaelbreno/go-bot/internal"
	"go.uber.org/zap"
)

var (
	connType   *string
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

	// Loading .env file
	if err := godotenv.Load(); err != nil {
		logger.Error(err.Error())
		os.Exit(0)
	}
	connType = flag.String("conn", "IRC", "Connection type to Twitch's Chat (IRC ou WS)")

	ctx = &internal.Context{
		Logger: logger,
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
