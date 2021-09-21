package main

import (
	"fmt"

	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/reader"
	"github.com/rafaelbreno/go-bot/services/message-reader/server"
	"github.com/rafaelbreno/go-bot/services/message-reader/storage"
)

func main() {
	ctx := internal.NewContext()

	ctx.Logger.Info("Starting service...")

	rd := reader.NewReader(ctx, storage.NewRedis(ctx))

	rd.Start()

	sv := server.NewServer(ctx)

	sv.Start()

	sv.Close()

	fmt.Println("Hello, World!")
}
