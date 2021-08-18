package conn

import (
	"fmt"
	"strings"
)

const (
	invalidConn = "The connection \"%s\" is invalid"
)

// Conn interface represents
// all action related to the
// connections types
type Conn interface {
	Listen()
}

// NewConn return a Conn pointer
func NewConn(connType *string) (Conn, error) {
	switch strings.ToUpper(*connType) {
	case "IRC":
		return Conn(&IRC{}), nil
	default:
		return *new(Conn), fmt.Errorf(invalidConn, *connType)
	}
}

// Listen fetch each message
func Listen(c Conn) {
	c.Listen()
}
