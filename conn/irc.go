package conn

import (
	"fmt"
	"net"
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
