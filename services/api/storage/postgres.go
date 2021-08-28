package storage

import (
	"fmt"
	"os"

	"github.com/rafaelbreno/go-bot/api/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB handler DB connection
type DB struct {
	Client *gorm.DB
	Ctx    *internal.Context
}

func newDB(ctx *internal.Context) *DB {
	var pgsqlURL string
	url := `host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`

	pgsqlURL = fmt.Sprintf(
		url,
		ctx.Env["PGSQL_HOST"],
		ctx.Env["PGSQL_PORT"],
		ctx.Env["PGSQL_USER"],
		ctx.Env["PGSQL_PASSWORD"],
		ctx.Env["PGSQL_DBNAME"],
	)
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
