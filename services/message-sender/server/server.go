package server

import (
	"context"

	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/proto"
	"github.com/rafaelbreno/go-bot/services/message-reader/repository"
)

// Server handles gRPC connection
type Server struct {
	proto.UnimplementedSenderServer
	Ctx  *internal.Context
	Repo *repository.MessageRepoCtx
}

// SendMessage receives a message body to treat
// and send to a channel's chat
func (s *Server) SendMessage(ctx context.Context, msg *proto.MessageRequest) *proto.Empty {
	return s.Repo.SendMessage(msg)
}
