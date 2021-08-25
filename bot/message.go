package bot

import "github.com/rafaelbreno/go-bot/internal"

// MsgType defines which message
// type was sent
type MsgType int

const (
	// Nil takes no action
	Nil MsgType = iota
	// Twitch 's communications
	Twitch
	// Ping to shakehands with Twitch
	Ping
	// Command is prefixed by exclamation mark !
	Command
)

// Message stores all information related
// to a sent message
type Message struct {
	Ctx         *internal.Context
	Type        MsgType
	SentBy      string
	SentMessage string
}

func (m *Message) getString() string {
	return b.Command.GetAnswer(m.SentBy, m.SentMessage)
}
