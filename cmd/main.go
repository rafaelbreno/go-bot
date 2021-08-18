package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rafaelbreno/go-bot/conn"
	"github.com/rafaelbreno/go-bot/internal"
	"go.uber.org/zap"
)

var (
	connType *string
	ctx      *internal.Context
	logger   *zap.Logger
)

func newLogger() {
	logger, _ = zap.NewProduction()
}

func newConnection() conn.Conn {
	c, err := conn.NewConn(connType)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(0)
	}

	return c
}

func init() {
	connType = flag.String("conn", "IRC", "Connection type to Twitch's Chat (IRC ou WS)")

	ctx = &internal.Context{
		Logger: logger,
		Conn:   newConnection(),
	}
}

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Loading .env file
	godotenv.Load()

	<-stop
}
