package storage

import "go.uber.org/zap"

type Postgres struct{}

func NewPostgres(l *zap.Logger) *Postgres {
	p := Postgres{}

	return &p
}
