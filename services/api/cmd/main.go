package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rafaelbreno/go-bot/api/config"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/server"
)

var ctx *internal.Context
var sv *server.Server

func init() {
	ctx = config.Ctx
	sv = server.NewServer()
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
