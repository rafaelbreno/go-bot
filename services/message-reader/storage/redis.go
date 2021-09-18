package storage

import (
	"context"
	"encoding/json"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/rafaelbreno/go-bot/services/message-reader/command"
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
)

type Redis struct {
	Conn *redis.Client
	Ctx  *internal.Context
}

// NewRedis create a new Redis instance
func NewRedis(ctx *internal.Context) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ctx.Env["REDIS_HOST"], ctx.Env["REDIS_PORT"]),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Redis{
		Conn: client,
		Ctx:  ctx,
	}
}

// GetChannels retrieves from Redis a list
// of all channels
func (r *Redis) GetChannels(key string) []string {
	var channels struct {
		Chs []string `json:"channels"`
	}

	res, err := r.Conn.Get(context.Background(), key).Result()

	if err != nil {
		r.Ctx.Logger.Error(err.Error())
		return channels.Chs
	}

	err = json.Unmarshal([]byte(res), &channels)

	if err != nil {
		r.Ctx.Logger.Error(err.Error())
		return channels.Chs
	}

	return channels.Chs
}

// GetCommand retrieves from Redis a command
// from a given channel and key
func (r *Redis) GetCommand(key string) command.Command {
	var cmd command.Command

	res, err := r.Conn.Get(context.Background(), key).Result()

	if err != nil {
		r.Ctx.Logger.Error(err.Error())
		return cmd
	}

	err = json.Unmarshal([]byte(res), &cmd)

	if err != nil {
		r.Ctx.Logger.Error(err.Error())
		return cmd
	}

	return cmd
}
