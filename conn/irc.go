package conn

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"

	"github.com/rafaelbreno/go-bot/internal"
)

const (
	ircConnURL = `%s:%s`
)

// IRC stores the information
// and actions related to IRC
// connection
type IRC struct {
	conn *net.Conn
	ctx  *internal.Context
}

// NewIRC returns a IRC
// struct pointer with
// configured connection
func NewIRC(ctx *internal.Context) (*IRC, error) {
	connStr := fmt.Sprintf(ircConnURL, os.Getenv("IRC_URL"), os.Getenv("IRC_PORT"))

	fmt.Println(connStr)

	conn, err := net.Dial("tcp", connStr)

	if err != nil {
		return &IRC{}, err
	}

	return &IRC{
		conn: &conn,
		ctx:  ctx,
	}, nil
}

// Listen start listen IRC
// channel
func (i *IRC) Listen() {
	tp := textproto.NewReader(bufio.NewReader(*i.conn))

	go func() {
		for {
			status, err := tp.ReadLine()
			if err != nil {
				i.ctx.Logger.Error(err.Error())
			}

			fmt.Println(status)
		}
	}()
}

// GetConn start listen IRC
// channel
func (i *IRC) GetConn() *net.Conn {
	return i.conn
}

// Ctx start listen IRC
// channel
func (i *IRC) Ctx() *internal.Context {
	return i.ctx
}
