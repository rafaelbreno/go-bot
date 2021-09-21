package repository

import (
	"github.com/rafaelbreno/go-bot/services/message-reader/conn"
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/proto"
	"github.com/rafaelbreno/go-bot/services/message-reader/reader"
	"github.com/rafaelbreno/go-bot/services/message-reader/sender"
)

type MessageRepo interface {
	SendMessage(msg *proto.MessageRequest) *proto.Empty
}

// MessageRepoCtx handles actions
// related to messages
type MessageRepoCtx struct {
	Ctx    *internal.Context
	Reader *reader.Reader
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

	m.Reader..SendMessage(msg.Channel, msg.Msg)

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
