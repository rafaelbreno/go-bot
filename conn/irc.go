package conn

import (
	"fmt"
	"net"
	"os"
)

const (
	ircConnURL = `%s:%s`
)

// IRC stores the information
// and actions related to IRC
// connection
type IRC struct {
	conn *net.Conn
}

// NewIRC returns a IRC
// struct pointer with
// configured connection
func NewIRC() (*IRC, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf(ircConnURL, os.Getenv("IRC_URL"), os.Getenv("IRC_PORT")))

	if err != nil {
		return &IRC{}, err
	}

	return &IRC{
		conn: &conn,
	}, nil
}
