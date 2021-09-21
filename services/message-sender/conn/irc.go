package conn

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"
	"time"

	"github.com/rafaelbreno/go-bot/services/message-sender/internal"
)

type IRC struct {
	Conn net.Conn
	Ctx  *internal.Context
}

const (
	ircConnURL = `%s:%s`
)

func NewIRC(ctx *internal.Context) *IRC {
	i := IRC{
		Ctx: ctx,
	}
	i.connect()
	i.pong()
	return &i
}

func (i *IRC) pong() {
	tp := textproto.NewReader(bufio.NewReader((i.Conn)))

	for {
		msg, err := tp.ReadLine()
		if err != nil {
			i.Close()
			i.connect()
		}
		if msg[:4] == "PING" {
			fmt.Fprint(i.Conn, "PONG")
			i.Ctx.Logger.Info("PONG")
		}
	}
}

func (i *IRC) connect() {
	connStr := fmt.Sprintf(ircConnURL, i.Ctx.Env["IRC_URL"], i.Ctx.Env["IRC_PORT"])

	var c net.Conn
	var err error

	connected := false

	for tries := 1; tries <= 3; tries++ {
		c, err = net.Dial("tcp", connStr)
		if err == nil {
			i.Conn = c
			connected = true
			break
		}
		errMsg := fmt.Sprintf("Error %s. Try number %d!", err.Error(), tries)
		i.Ctx.Logger.Error(errMsg)
		time.Sleep(2 * time.Second)
	}

	if !connected {
		i.Ctx.Logger.Error("Unable to connect to IRC")
		os.Exit(0)
	}

	pass := fmt.Sprintf("PASS %s\r\n", i.Ctx.Env["BOT_OAUTH_TOKEN"])
	if _, err := fmt.Fprint(i.Conn, pass); err != nil {
		i.Ctx.Logger.Error(err.Error())
		os.Exit(0)
	}

	nick := fmt.Sprintf("NICK %s\r\n", i.Ctx.Env["BOT_NAME"])
	if _, err := fmt.Fprint(i.Conn, nick); err != nil {
		i.Ctx.Logger.Error(err.Error())
		os.Exit(0)
	}
	i.Ctx.Logger.Info(nick)
}

// Close ends IRC connection
func (i *IRC) Close() {
	i.Ctx.Logger.Info("Closing IRC connection")
	if err := i.Conn.Close(); err != nil {
		i.Ctx.Logger.Error(err.Error())
		os.Exit(0)
	}
	i.Ctx.Logger.Info("IRC connection closed")
}
