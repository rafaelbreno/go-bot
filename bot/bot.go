package bot

import (
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

// Start ignites the bot
func Start(ctx *internal.Context, irc *conn.IRC) {
	b = &Bootstrap{
		Ctx: ctx,
		IRC: irc,
	}

	go b.IRC.Listen()
}
