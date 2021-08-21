package bot

import (
	"fmt"

	"github.com/rafaelbreno/go-bot/command"
	"github.com/rafaelbreno/go-bot/conn"
	"github.com/rafaelbreno/go-bot/internal"
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
	b.ParseMsg()
}

// ParseMsg aa
func (b *Bootstrap) ParseMsg() {
	b.Ctx.Logger.Info("Start parser")
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}
}
