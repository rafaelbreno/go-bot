package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/rafaelbreno/go-bot/api/internal"
)

type InMem struct {
	Ctx    *internal.Context
	Client *redis.Client
}

func newInMem(ctx *internal.Context) *InMem {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	cmd := rdb.Ping(context.Background())

	if err := cmd.Err(); err != nil {
		ctx.Logger.Error(err.Error())
		os.Exit(0)
	}

	return &InMem{
		Ctx:    ctx,
		Client: rdb,
	}
}
