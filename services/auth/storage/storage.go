package storage

import "go.uber.org/zap"

type Storage struct {
	Redis *Redis
	Pg    *Postgres
}

func NewStorage(l *zap.Logger) *Storage {
	s := Storage{
		Redis: NewRedis(l),
		Pg:    NewPostgres(l),
	}

	return &s
}
