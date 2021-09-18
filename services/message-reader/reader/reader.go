package reader

import (
	"github.com/rafaelbreno/go-bot/services/message-reader/conn"
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/storage"
)

type Reader struct {
	Storage storage.Storage
	IRC     *conn.IRC
	Ctx     *internal.Context
}

func NewReader(ctx *internal.Context, st storage.Storage) *Reader {
	irc := conn.NewIRC(ctx)

	return &Reader{
		Storage: st,
		Ctx:     ctx,
		IRC:     irc,
	}
}

func (r *Reader) Start() {
	for {
		select {
		case msg := <-r.IRC.Msg:
			r.Ctx.Logger.Info(msg)
		}
	}
}
