package bot

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rafaelbreno/go-bot/internal"
)

type Parser struct {
	Ctx *internal.Context
	// UserRegex retrieves username
	UserRegex *regexp.Regexp

	// MessageRegex retrieves user's message
	MessageRegex *regexp.Regexp

	channelPrefix string
}

func NewParser(ctx *internal.Context) *Parser {
	msgRegexStr := fmt.Sprintf(`(#%s :).{1,}$`, ctx.ChannelName)
	ctx.Logger.Info("Initialzied parser")
	return &Parser{
		Ctx:           ctx,
		UserRegex:     regexp.MustCompile(`^(:)[a-zA-Z0-9_]{1,}(!)`),
		MessageRegex:  regexp.MustCompile(msgRegexStr),
		channelPrefix: fmt.Sprintf("#%s :", ctx.ChannelName),
	}
}

// ParseMsg a string into a
// Message struct
func (p *Parser) ParseMsg(msgStr string) *Message {
	fmt.Println(msgStr)

	if strings.HasPrefix(msgStr, "PING") {
		p.Ctx.Logger.Info("Received ping")
		return &Message{
			Type: Ping,
		}
	}
	if strings.HasPrefix(msgStr, ":tmi.twitch.tv") {
		return &Message{
			Type: Nil,
		}
	}

	sentMessageRaw := p.MessageRegex.FindString(msgStr)
	sentByRaw := p.UserRegex.Find([]byte(msgStr))
	lenSentBy := len(sentByRaw)

	if lenSentBy == 0 {
		p.Ctx.Logger.Info("lenSentBy = 0")
		return &Message{
			Type: Nil,
		}
	}

	sentMessage := strings.TrimPrefix(sentMessageRaw, p.channelPrefix)
	sentBy := string(sentByRaw[1 : lenSentBy-1])

	if !strings.HasPrefix(sentMessage, "!") {
		return &Message{
			Type: Nil,
		}
	}

	return &Message{
		Type:        Command,
		SentMessage: sentMessage,
		SentBy:      sentBy,
	}
}
