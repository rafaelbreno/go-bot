package storage

import (
	"github.com/go-redis/redis/v8"
	"github.com/rafaelbreno/go-bot/auth/internal"
)

type Redis struct {
	common *internal.Common
	Conn   *redis.Client
}

func NewRedis(c *internal.Common) *Redis {
	r := Redis{
		common: c,
	}

	r.setConnection()

	return &r
}

func (r *Redis) setConnection() {
	r.Conn = redis.NewClient(&redis.Options{
		Addr:     r.getAddress(),
		Password: r.common.Env.RedisPassword,
		DB:       r.common.Env.RedisDB,
	})
}

func (r *Redis) getAddress() string {
	if r.common.Env.RedisPort == "" {
		return r.common.Env.RedisHost
	}

	return string(r.common.Env.RedisHost + ":" + r.common.Env.RedisPort)
}
