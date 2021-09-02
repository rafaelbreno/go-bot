package storage

import (
	"github.com/rafaelbreno/go-bot/api/internal"
)

// Storage manages connections between
// this API and any resource of storage
// e.g(Postgres, S3, etc.).
type Storage struct {
	SQL         *DB
	KafkaClient *KafkaClient
	Ctx         *internal.Context
}

// NewStorage return a storage manager
// for database and in-memory
func NewStorage(ctx *internal.Context) *Storage {
	return &Storage{
		Ctx:         ctx,
		SQL:         newDB(ctx),
		KafkaClient: newKafkaClient(ctx),
	}
}
