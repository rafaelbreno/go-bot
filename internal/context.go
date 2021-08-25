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

// WriteContexts returns multiples contexts,
// each for one different channels
func WriteContexts(l *zap.Logger, authToken, botName string, channels []string) map[string]*Context {
	var chs map[string]*Context

	for _, channel := range channels {
		if _, ok := chs[channel]; !ok {
			chs[channel] = &Context{
				Logger:      l,
				ChannelName: channel,
				OAuthToken:  authToken,
				BotName:     botName,
			}
		}
	}

	return chs
}
