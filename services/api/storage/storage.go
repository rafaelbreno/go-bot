package storage

import (
	"github.com/rafaelbreno/go-bot/api/internal"
)

type Storage struct {
	SQL   *DB
	InMem *InMem
	Ctx   *internal.Context
}

// NewStorage return a storage manager
// for database and in-memory
func NewStorage(ctx *internal.Context) *Storage {
	return &Storage{
		Ctx:   ctx,
		SQL:   newDB(ctx),
		InMem: newInMem(ctx),
	}
}
