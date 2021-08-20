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
	conn net.Conn
	ctx  *internal.Context
	tp   *textproto.Reader
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
		conn: c,
		ctx:  ctx,
	}, err
}

// Listen start listen IRC
// channel
func (i *IRC) Listen() {
	i.connect()

	i.tp = textproto.NewReader(bufio.NewReader(i.conn))

	go func() {
		for {
			msg, err := i.tp.ReadLine()
			if err != nil {
				i.ctx.Logger.Error(err.Error())
				os.Exit(0)
			}
			fmt.Println(msg)
		}
	}()
}

func (i *IRC) connect() {
	fmt.Fprintf(i.conn, "PASS %s\r\n", os.Getenv("BOT_OAUTH_TOKEN"))
	fmt.Fprintf(i.conn, "NICK %s\r\n", os.Getenv("BOT_USERNAME"))
	fmt.Fprintf(i.conn, "JOIN #%s\r\n", os.Getenv("CHANNEL_NAME"))
}

// GetConn start listen IRC
// channel
func (i *IRC) GetConn() net.Conn {
	return i.conn
}

// Ctx start listen IRC
// channel
func (i *IRC) Ctx() *internal.Context {
	return i.ctx
}

// Close ends IRC connection
func (i *IRC) Close() {
	i.ctx.Logger.Info("Closing IRC connection")
	if err := i.conn.Close(); err != nil {
		i.ctx.Logger.Error(err.Error())
		os.Exit(0)
	}
	i.ctx.Logger.Info("IRC connection closed")
}
