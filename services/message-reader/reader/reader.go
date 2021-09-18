package reader

import (
	"fmt"

	"github.com/rafaelbreno/go-bot/services/message-reader/conn"
	"github.com/rafaelbreno/go-bot/services/message-reader/helpers"
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/message"
	"github.com/rafaelbreno/go-bot/services/message-reader/storage"
)

type Reader struct {
	Channels []string
	Storage  storage.Storage
	IRC      *conn.IRC
	Ctx      *internal.Context
}

func NewReader(ctx *internal.Context, st storage.Storage) *Reader {
	irc := conn.NewIRC(ctx)

	return &Reader{
		Channels: []string{},
		Storage:  st,
		Ctx:      ctx,
		IRC:      irc,
	}
}

func (r *Reader) Start() {
	r.joinRedis()
	p := message.Parser{
		Ctx: r.Ctx,
		IRC: r.IRC,
	}

	for {
		select {
		case msg := <-r.IRC.Msg:
			go p.Send(msg)
		}
	}
}

func (r *Reader) joinRedis() {
	for _, c := range storage.GetChannels(r.Storage) {
		r.Channels = append(r.Channels, c)
		r.JoinChat(c)
	}
}

func (r *Reader) CheckChannel(name string) {
	if !helpers.FindInSliceStr(r.Channels, name) {
		r.Channels = append(r.Channels, name)
		r.JoinChat(name)
	}
}

func (r *Reader) JoinChat(name string) {
	r.write(fmt.Sprintf("JOIN #%s", name))
}

func (r *Reader) partChat(name string) {
	r.write(fmt.Sprintf("PART #%s", name))
	r.Channels = helpers.RemoveElementStr(r.Channels, name)
}

func (r *Reader) write(msg string) {
	r.Ctx.Logger.Info("Sending Message")
	_, err := r.IRC.Conn.Write([]byte(fmt.Sprintf("%s\r\n", msg)))
	if err != nil {
		r.Ctx.Logger.Error(err.Error())
	}
}
