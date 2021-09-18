package message

import (
	"regexp"

	"github.com/rafaelbreno/go-bot/services/message-reader/command"
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/storage"
)

type Message struct {
	Channel string
	SentBy  string
	Value   string
}

var (
	triggerRegex = regexp.MustCompile(`![a-zA-Z0-9]{1,}`)
)

var (
	nilCommand = command.Command{}
)

func (m *Message) Send(ctx *internal.Context, st *storage.Storage) {
	trigger := triggerRegex.FindString(m.Value)

	cmd := storage.GetCommand(m.Channel, trigger, *st)

	if cmd == nilCommand {
		return
	}

	ans := cmd.Parse()
}
