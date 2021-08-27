package storage

import (
	"fmt"
	"os"

	"github.com/rafaelbreno/go-bot/api/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgsqlURL string

// DB handler DB connection
type DB struct {
	Client *gorm.DB
	Ctx    *internal.Context
}

func init() {
	url := `host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`

	pgsqlURL = fmt.Sprintf(
		url,
		os.Getenv("PGSQL_HOST"),
		os.Getenv("PGSQL_PORT"),
		os.Getenv("PGSQL_USER"),
		os.Getenv("PGSQL_PASSWORD"),
		os.Getenv("PGSQL_DBNAME"),
	)
}

func newDB(ctx *internal.Context) *DB {
	db, err := gorm.Open(postgres.Open(pgsqlURL), &gorm.Config{})

	if err != nil {
		ctx.Logger.Error(err.Error())
		os.Exit(0)
	}

	return &DB{
		Ctx:    ctx,
		Client: db,
	}
}
