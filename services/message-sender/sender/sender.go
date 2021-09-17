package sender

import (
	"fmt"

	"github.com/rafaelbreno/go-bot/services/message-reader/conn"
	"github.com/rafaelbreno/go-bot/services/message-reader/helpers"
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
)

// Sender manages the channel that are connected
type Sender struct {
	Channels []string
	Ctx      *internal.Context
	IRC      *conn.IRC
}

// NewSender creates an instance for Sender struct
func NewSender(ctx *internal.Context, irc *conn.IRC) *Sender {
	return &Sender{
		Channels: []string{},
		Ctx:      ctx,
		IRC:      irc,
	}
}

// CheckChannel verify if there's a channel already connected
func (s *Sender) CheckChannel(name string) {
	if !helpers.FindInSliceStr(s.Channels, name) {
		s.Channels = append(s.Channels, name)
		s.joinChat(name)
	}
}

// SendMessage
func (s *Sender) SendMessage(channel, message string) {
	s.CheckChannel(channel)
	s.write(fmt.Sprintf("PRIVMSG #%s :%s", channel, message))
}

func (s *Sender) joinChat(name string) {
	s.write(fmt.Sprintf("JOIN #%s", name))
}

func (s *Sender) partChat(name string) {
	s.write(fmt.Sprintf("PART #%s", name))
	s.Channels = helpers.RemoveElementStr(s.Channels, name)
}

func (s *Sender) write(msg string) {
	s.Ctx.Logger.Info("Sending Message")
	_, err := s.IRC.Conn.Write([]byte(fmt.Sprintf("%s\r\n", msg)))
	if err != nil {
		s.Ctx.Logger.Error(err.Error())
	}
}
