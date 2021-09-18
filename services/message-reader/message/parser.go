package message

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rafaelbreno/go-bot/services/message-reader/conn"
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
)

var (
	channelNameRegex = regexp.MustCompile(`#[a-zA-Z0-9\_]{1,}`)
)

type Parser struct {
	Ctx *internal.Context
	IRC *conn.IRC
}

func (p *Parser) Pong() {
	fmt.Fprint(p.IRC.Conn, fmt.Sprint("PONG\r\n"))
}

var (
	channelRegex = regexp.MustCompile(`#[a-zA-Z0-9\_]{1,}`)
)

// :ricardinst!ricardinst@ricardinst.tmi.twitch.tv PRIVMSG #rafiusky :e
func (p *Parser) Parse(msg string) *Message {
	fmt.Println(msg)

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

	if len(val) <= 3 {
		return &Message{}
	}

	return &Message{
		SentBy:  string(sentBy),
		Channel: channel[1:],
		Value:   val[2],
	}
}
