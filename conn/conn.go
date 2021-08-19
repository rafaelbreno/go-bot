package conn

import (
	"fmt"
	"net"
	"strings"

	"github.com/rafaelbreno/go-bot/internal"
)

const (
	invalidConn = "The connection \"%s\" is invalid"
)

// Conn interface represents
// all action related to the
// connections types
type Conn interface {
	Listen()
	GetConn() *net.Conn
	Ctx() *internal.Context
}

// NewConn return a Conn pointer
func NewConn(connType *string, ctx *internal.Context) (Conn, error) {
	switch strings.ToUpper(*connType) {
	case "IRC":
		irc, err := NewIRC(ctx)
		return Conn(irc), err
	default:
		return *new(Conn), fmt.Errorf(invalidConn, *connType)
	}
}

// Listen fetch each message
func Listen(c Conn) {
	c.Ctx().Logger.Info("Start listening to Twitch's Chat")
	c.Listen()
}

// GetConn a
func GetConn(c Conn) *net.Conn {
	return c.GetConn()
}
