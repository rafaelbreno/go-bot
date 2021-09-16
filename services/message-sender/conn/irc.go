package conn

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"
	"time"

	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
)

type IRC struct {
	Conn net.Conn
	Ctx  *internal.Context
	TP   *textproto.Reader
	Msg  chan string
}

const (
	ircConnURL = `%s:%s`
)

func (i *IRC) Listen(ch chan string) {
	i.connect()

	go func() {
		for {
			msg, err := i.TP.ReadLine()
			if err != nil {
				close(i.Msg)
				i.connect()
				continue
			}
			ch <- msg
		}
	}()
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

	pass := fmt.Sprintf("PASS %s\r\n", i.Ctx.OAuthToken)
	if _, err := fmt.Fprint(i.Conn, pass); err != nil {
		i.Ctx.Logger.Error(err.Error())
		os.Exit(0)
	}

	nick := fmt.Sprintf("NICK %s\r\n", i.Ctx.BotName)
	if _, err := fmt.Fprint(i.Conn, nick); err != nil {
		i.Ctx.Logger.Error(err.Error())
		os.Exit(0)
	}
	i.Ctx.Logger.Info(nick)
	join := fmt.Sprintf("JOIN #%s\r\n", i.Ctx.ChannelName)
	if _, err := fmt.Fprint(i.Conn, join); err != nil {
		i.Ctx.Logger.Error(err.Error())
		os.Exit(0)
	}
	i.Ctx.Logger.Info(join)

	i.TP = textproto.NewReader(bufio.NewReader(i.Conn))
	i.Msg = make(chan string, 1)
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
