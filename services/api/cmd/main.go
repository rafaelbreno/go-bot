package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rafaelbreno/go-bot/api/internal"
	"go.uber.org/zap"
)

var ctx *internal.Context

func init() {
	l, err := zap.NewProduction()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	ctx = &internal.Context{
		Logger: l,
	}
}

func main() {
	ctx.Logger.Info("Starting app")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	ctx.Logger.Info("Gracefully terminating...")
}
