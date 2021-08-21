package bot

import (
	"fmt"

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
	Command *command.CommandCtx
}

var b *Bootstrap
var ch chan string

// Start ignites the bot
func Start(ctx *internal.Context, irc *conn.IRC) {
	ch = make(chan string, 1)

	b = &Bootstrap{
		Ctx: ctx,
		IRC: irc,
		Command: &command.CommandCtx{
			Ctx: ctx,
		},
	}
	b.Ctx.Logger.Info("Start bot")

	go b.IRC.Listen(ch)
	b.receiveMsg()
}

func (b *Bootstrap) receiveMsg() {
	b.Ctx.Logger.Info("Start parser")
	p := NewParser(b.Ctx)
	for {
		select {
		case msgStr := <-ch:
			msg := p.ParseMsg(msgStr)
			msg.Ctx = b.Ctx
			b.do(msg)
		}
	}
}

func (b *Bootstrap) do(msg *Message) {
	switch msg.Type {
	case Nil:
		break
	case Ping:
		utils.Write(b.Ctx, b.IRC.Conn, "PONG")
		break
	case Command:
		utils.Write(b.Ctx, b.IRC.Conn, fmt.Sprintf("PRIVMSG #%s :%s", b.Ctx.ChannelName, msg.getString()))
		break
	}
}
