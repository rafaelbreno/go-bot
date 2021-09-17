package repository

import (
	"fmt"

	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/proto"
)

type MessageRepo interface {
	SendMessage(msg *proto.MessageRequest) *proto.Empty
}

// MessageRepoCtx handles actions
// related to messages
type MessageRepoCtx struct {
	Ctx *internal.Context
}

// SendMessage receives a message
// and send it to a channel's chat
func (m *MessageRepoCtx) SendMessage(msg *proto.MessageRequest) *proto.Empty {
	fmt.Printf("%s: %s\n", msg.Channel, msg.Msg)

	return &proto.Empty{}
}

// NewMessageRepo builds a message repository
// with given context
func NewMessageRepo(ctx *internal.Context) *MessageRepoCtx {
	return &MessageRepoCtx{
		Ctx: ctx,
	}
}
