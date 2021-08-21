package bot

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rafaelbreno/go-bot/internal"
)

// MsgType defines which message
// type was sent
type MsgType int

const (
	// Nil takes no action
	Nil MsgType = iota
	// Twitch 's communications
	Twitch
	// User is the common user
	User
	// VIP is the user vip
	VIP
	// MOD is the moderator
	MOD
	// Streamer is the streamer
	Streamer
	// Ping to shakehands with Twitch
	Ping
)

type Parser struct {
	Ctx *internal.Context
	// UserRegex retrieves username
	UserRegex *regexp.Regexp

	// MessageRegex retrieves user's message
	MessageRegex *regexp.Regexp
}

func NewParser(ctx *internal.Context) *Parser {
	msgRegexStr := fmt.Sprintf(`(#%s :).{1,}$`, ctx.ChannelName)
	return &Parser{
		Ctx:          ctx,
		UserRegex:    regexp.MustCompile(`^(:)[a-zA-Z0-9]{1,}(!)`),
		MessageRegex: regexp.MustCompile(msgRegexStr),
	}
}

// Message stores all information related
// to a sent message
type Message struct {
	Type        MsgType
	SentBy      string
	SentMessage string
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

	sentMessage := p.MessageRegex.FindString(msgStr)
	sentBy := p.UserRegex.Find([]byte(msgStr))
	lenSentBy := len(sentBy)

	if lenSentBy == 0 {
		return &Message{
			Type: Nil,
		}
	}

	return &Message{
		SentMessage: strings.TrimPrefix(sentMessage, "#rafiusky :"),
		SentBy:      string(sentBy[1 : lenSentBy-1]),
	}
}
