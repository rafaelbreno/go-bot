package storage

import (
	pg "github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/rafaelbreno/go-bot/auth/internal"
	"github.com/rafaelbreno/go-bot/auth/user"
)

type Postgres struct {
	common *internal.Common
	Conn   *pg.DB
}

func NewPostgres(c *internal.Common) *Postgres {
	p := Postgres{
		common: c,
	}

	p.
		setConnection().
		migration()

	return &p
}

func (p *Postgres) setConnection() *Postgres {
	p.Conn = pg.Connect(&pg.Options{
		Addr:     p.getAddress(),
		User:     p.common.Env.PgUser,
		Password: p.common.Env.PgPassword,
		Database: p.common.Env.PgDBName,
	})
	return p
}

func (p *Postgres) getAddress() string {
	if p.common.Env.PgPort == "" {
		return p.common.Env.PgHost
	}
	return string(p.common.Env.PgHost + ":" + p.common.Env.PgPort)
}

func (p *Postgres) migration() {
	models := []interface{}{
		(*user.User)(nil),
	}

	for _, m := range models {
		if err := p.Conn.Model(m).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		}); err != nil {
			p.common.Logger.Error(err.Error())
		}
	}
}
