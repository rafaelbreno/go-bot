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
	// Twitch 's communications
	Twitch MsgType = iota
	// User is the common user
	User
	// VIP is the user vip
	VIP
	// MOD is the moderator
	MOD
	// Streamer is the streamer
	Streamer
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
	RawMessage  string
	Type        MsgType
	SentBy      string
	SentMessage string
}

// ParseMsg a string into a
// Message struct
func (p *Parser) ParseMsg(msgStr string) {
	fmt.Println(msgStr)

	sentMessage := p.MessageRegex.FindString(msgStr)
	sentBy := p.UserRegex.Find([]byte(msgStr))
	lenSentBy := len(sentBy)

	if lenSentBy == 0 {
		return
	}

	msg := Message{
		RawMessage:  msgStr,
		SentMessage: strings.TrimPrefix(sentMessage, "#rafiusky :"),
		SentBy:      string(sentBy[1 : lenSentBy-1]),
	}

	fmt.Printf("SentBy %s\n", msg.SentBy)
	fmt.Printf("SentMessage %s\n", msg.SentMessage)
}
