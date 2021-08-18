package internal

import (
	"github.com/rafaelbreno/go-bot/conn"
	"go.uber.org/zap"
)

// Context stores informations related
// to the connection
type Context struct {
	// Logger access a logging package
	// to log all action inside the bot
	Logger *zap.Logger
	// Conn is the connection between the
	// bot and Twitch's chat
	Conn *conn.Conn
}
