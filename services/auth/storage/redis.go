package storage

import "go.uber.org/zap"

type Redis struct{}

func NewRedis(l *zap.Logger) *Redis {
	r := Redis{}

	return &r
}
