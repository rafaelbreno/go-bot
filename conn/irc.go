package conn

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"
	"time"

	"github.com/rafaelbreno/go-bot/internal"
)

const (
	ircConnURL = `%s:%s`
)

// IRC stores the information
// and actions related to IRC
// connection
type IRC struct {
	Conn net.Conn
	Ctx  *internal.Context
	TP   *textproto.Reader
}

// NewIRC returns a IRC
// struct pointer with
// configured connection
func NewIRC(ctx *internal.Context) (*IRC, error) {
	connStr := fmt.Sprintf(ircConnURL, os.Getenv("IRC_URL"), os.Getenv("IRC_PORT"))

	fmt.Println(connStr)

	var c net.Conn
	var err error

	for tries := 1; tries <= 3; tries++ {
		c, err = net.Dial("tcp", connStr)
		if err == nil {
			break
		}
		errMsg := fmt.Sprintf("Error %s. Try number %d!", err.Error(), tries)
		ctx.Logger.Error(errMsg)
		c.Close()
		time.Sleep(2 * time.Second)
	}

	return &IRC{
		Conn: c,
		Ctx:  ctx,
	}, err
}

// Listen start listen IRC channel,
// and if it disconnects, it will try
// to reconnect 3 times
func (i *IRC) Listen() {
	i.connect()

	go func() {
		for {
			msg, err := i.TP.ReadLine()
			if err != nil {
				i.Ctx.Logger.Error(err.Error())
				i.connect()
			}
			fmt.Println(msg)
		}
	}()
}

func (i *IRC) connect() {
	fmt.Fprintf(i.Conn, "PASS %s\r\n", os.Getenv("BOT_OAUTH_TOKEN"))
	fmt.Fprintf(i.Conn, "NICK %s\r\n", os.Getenv("BOT_USERNAME"))
	fmt.Fprintf(i.Conn, "JOIN #%s\r\n", os.Getenv("CHANNEL_NAME"))

	i.TP = textproto.NewReader(bufio.NewReader(i.Conn))
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
