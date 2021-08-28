package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rafaelbreno/go-bot/api/entity"
	"github.com/rafaelbreno/go-bot/api/handler"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/server"
	"github.com/rafaelbreno/go-bot/api/storage"
	"go.uber.org/zap"
)

var ctx *internal.Context
var sv *server.Server
var st *storage.Storage
var h handler.Handler

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
		Env: map[string]string{
			"PGSQL_HOST":     os.Getenv("PGSQL_HOST"),
			"PGSQL_PORT":     os.Getenv("PGSQL_PORT"),
			"PGSQL_USER":     os.Getenv("PGSQL_USER"),
			"PGSQL_PASSWORD": os.Getenv("PGSQL_PASSWORD"),
			"PGSQL_DBNAME":   os.Getenv("PGSQL_DBNAME"),
			"REDIS_HOST":     os.Getenv("REDIS_HOST"),
			"REDIS_PORT":     os.Getenv("REDIS_PORT"),
			"REDIS_PASSWORD": os.Getenv("REDIS_PASSWORD"),
		},
	}

	sv = &server.Server{
		Ctx:  ctx,
		Port: os.Getenv("API_PORT"),
	}

	st = storage.NewStorage(ctx)

	st.
		SQL.
		Client.AutoMigrate(&entity.Command{})

	h = handler.Handler{
		Ctx:     ctx,
		Storage: st,
	}
}

func main() {
	ctx.Logger.Info("Starting app")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go sv.ListenAndServe(h)

	defer sv.Close()
	<-stop

	ctx.Logger.Info("Gracefully terminating...")
}
