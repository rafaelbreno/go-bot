package message

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rafaelbreno/go-bot/services/message-reader/conn"
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
)

var (
	channelNameRegex = regexp.MustCompile(`#[a-zA-Z0-9_]+`)
)

type Parser struct {
	Ctx *internal.Context
	IRC *conn.IRC
}

func (p *Parser) Pong() {
	p.Ctx.Logger.Info("PONG")
	_, err := fmt.Fprint(p.IRC.Conn, "PONG\r\n")
	if err != nil {
		p.Ctx.Logger.Error(err.Error())
	}
}

// :ricardinst!ricardinst@ricardinst.tmi.twitch.tv PRIVMSG #rafiusky :e
func (p *Parser) Parse(msg string) *Message {
	p.Ctx.Logger.Info(msg)

	if msg[0:4] == "PING" {
		p.Pong()
		return &Message{}
	}

	byteMsg := []byte(msg)
	var sentBy []byte

	for _, b := range byteMsg {
		if b == byte('!') {
			break
		}
		sentBy = append(sentBy, b)
	}

	channel := channelNameRegex.FindString(msg)

	if channel == "" {
		return &Message{}
	}

	val := strings.SplitN(msg, ":", 3)

	if len(val) < 3 {
		return &Message{}
	}

	return &Message{
		SentBy:  string(sentBy[1:]),
		Channel: channel[1:],
		Value:   val[2],
	}
}
