package message

import (
	"regexp"

	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/proto"
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

func (m *Message) Send(ctx *internal.Context, st *storage.Storage) {
	trigger := triggerRegex.FindString(m.Value)

	if trigger == "" {
		return
	}

	cmd := storage.GetCommand(m.Channel, trigger, *st)

	*ctx.Msg <- &proto.MessageRequest{
		Channel: m.Channel,
		Msg:     cmd.Parse(),
	}
}
