package bot

import (
	"github.com/rafaelbreno/go-bot/command"
	"github.com/rafaelbreno/go-bot/conn"
	"github.com/rafaelbreno/go-bot/internal"
	"github.com/rafaelbreno/go-bot/utils"
)

// Bootstrap manages all actions
// related to the bot itself
type Bootstrap struct {
	Ctx     *internal.Context
	IRC     *conn.IRC
	Command *command.Command
}

var b *Bootstrap
var ch chan string

// Start ignites the bot
func Start(ctx *internal.Context, irc *conn.IRC) {
	ch = make(chan string, 1)

	b = &Bootstrap{
		Ctx: ctx,
		IRC: irc,
	}
	b.Ctx.Logger.Info("Start bot")

	go b.IRC.Listen(ch)
	b.ReceiveMsg()
}

// ReceiveMsg aa
func (b *Bootstrap) ReceiveMsg() {
	b.Ctx.Logger.Info("Start parser")
	p := NewParser(b.Ctx)
	for {
		select {
		case msgStr := <-ch:
			b.Do(p.ParseMsg(msgStr))
		}
	}
}

func (b *Bootstrap) Do(msg *Message) {
	switch msg.Type {
	case Nil:
		break
	case Ping:
		utils.Write(b.Ctx, b.IRC.Conn, "PONG")
		break
	}
}
