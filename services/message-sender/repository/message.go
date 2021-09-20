package repository

import (
	"github.com/rafaelbreno/go-bot/services/message-sender/conn"
	"github.com/rafaelbreno/go-bot/services/message-sender/internal"
	"github.com/rafaelbreno/go-bot/services/message-sender/proto"
	"github.com/rafaelbreno/go-bot/services/message-sender/sender"
)

type MessageRepo interface {
	SendMessage(msg *proto.MessageRequest) *proto.Empty
}

// MessageRepoCtx handles actions
// related to messages
type MessageRepoCtx struct {
	Ctx    *internal.Context
	Sender *sender.Sender
}

// SendMessage receives a message
// and send it to a channel's chat
func (m *MessageRepoCtx) SendMessage(msg *proto.MessageRequest) *proto.Empty {
	if msg.Msg == "" {
		m.Ctx.Logger.Error("empty 'msg' field")
		return &proto.Empty{}
	}
	if msg.Channel == "" {
		m.Ctx.Logger.Error("empty 'channel' field")
		return &proto.Empty{}
	}

	m.Sender.SendMessage(msg.Channel, msg.Msg)

	return &proto.Empty{}
}

// NewMessageRepo builds a message repository
// with given context
func NewMessageRepo(ctx *internal.Context) *MessageRepoCtx {
	return &MessageRepoCtx{
		Ctx:    ctx,
		Sender: sender.NewSender(ctx, conn.NewIRC(ctx)),
	}
}
