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
	Env         map[string]string
}

type Channel struct {
	Name string
	Env  map[string]string
}

// WriteContexts returns multiples contexts,
// each for one different channels
func WriteContexts(l *zap.Logger, authToken, botName string, channels []Channel) map[string]*Context {
	chs := map[string]*Context{}

	for _, channel := range channels {
		if _, ok := chs[channel.Name]; !ok {
			chs[channel.Name] = &Context{
				Logger:      l,
				ChannelName: channel.Name,
				OAuthToken:  authToken,
				BotName:     botName,
				Env:         channel.Env,
			}
		}
	}

	return chs
}
