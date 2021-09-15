package main

import (
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
	connection []*conn.IRC
	ctxs       map[string]*internal.Context
	logger     *zap.Logger
)

func newLogger() {
	logger, _ = zap.NewProduction()
}

func setConnections() {
	for _, ctx := range ctxs {
		c, err := conn.NewIRC(ctx)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(0)
		}

		connection = append(connection, c)
	}
}

func init() {
	newLogger()

	// Loading .env file
	if err := godotenv.Load(); err != nil {
		logger.Error(err.Error())
	}
	ctxs = internal.WriteContexts(logger,
		os.Getenv("BOT_OAUTH_TOKEN"),
		os.Getenv("BOT_USERNAME"),
		[]internal.Channel{
			{
				Name: "rafiusky",
				Env: map[string]string{
					"SPOTIFY_TOKEN": os.Getenv("SPOTIFY_TOKEN"),
				},
			},
		})

	setConnections()
}

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	for _, c := range connection {
		go bot.Start(c.Ctx, c)

		defer c.Close()
	}

	<-stop

	logger.Info("Gracefully terminating bot")
}
