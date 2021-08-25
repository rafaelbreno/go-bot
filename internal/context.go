package internal

import (
	"go.uber.org/zap"
)

// Context stores informations related
// to the connection
type Context struct {
	// Logger access a logging package
	// to log all action inside the bot
	Logger      *zap.Logger
	ChannelName string
	OAuthToken  string
	BotName     string
}
