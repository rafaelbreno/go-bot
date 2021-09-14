package main

import (
	"os"

	"github.com/rafaelbreno/go-bot/api/config"
	"github.com/rafaelbreno/go-bot/api/entity"
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
	config.Storage.SQL.Client.AutoMigrate(&entity.Command{})

	ctx.Logger.Info("Starting app")

	sv.ListenAndServe()

	defer sv.Close()

	st := config.Storage

	defer st.KafkaClient.P.Close()
	db, _ := st.SQL.Client.DB()
	defer func() {
		if err := db.Close(); err != nil {
			ctx.Logger.Error(err.Error())
			os.Exit(0)
		}
	}()
	ctx.Logger.Info("Gracefully terminating...")
}
