package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/server"
	"go.uber.org/zap"
)

var ctx *internal.Context
var sv *server.Server

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

	ctx = &internal.Context{
		Logger: l,
	}

	sv = &server.Server{
		Ctx:  ctx,
		Port: os.Getenv("API_PORT"),
	}
}

func main() {
	ctx.Logger.Info("Starting app")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go sv.ListenAndServe()

	defer sv.Close()
	<-stop

	ctx.Logger.Info("Gracefully terminating...")
}
