package storage

import (
	pg "github.com/go-pg/pg/v10"
	"github.com/rafaelbreno/go-bot/auth/internal"
)

type Postgres struct {
	common *internal.Common
	Conn   *pg.DB
}

func NewPostgres(c *internal.Common) *Postgres {
	p := Postgres{
		common: c,
	}

	p.setConnection()

	return &p
}

func (p *Postgres) setConnection() {
	p.Conn = pg.Connect(&pg.Options{
		Addr:     p.getAddress(),
		User:     p.common.Env.PgUser,
		Password: p.common.Env.PgPassword,
		Database: p.common.Env.PgDBName,
	})
}

func (p *Postgres) getAddress() string {
	if p.common.Env.PgPort == "" {
		return p.common.Env.PgHost
	}
	return string(p.common.Env.PgHost + ":" + p.common.Env.PgPort)
}
