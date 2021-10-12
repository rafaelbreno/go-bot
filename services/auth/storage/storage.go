package storage

import (
	"github.com/rafaelbreno/go-bot/auth/internal"
)

type Storage struct {
	Redis *Redis
	Pg    *Postgres
}

func NewStorage(c *internal.Common) *Storage {
	s := Storage{
		Redis: NewRedis(c),
		Pg:    NewPostgres(c),
	}

	return &s
}
