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
func (s *Server) SendMessage(ctx context.Context, msg *proto.MessageRequest) (*proto.Empty, error) {
	return s.Repo.SendMessage(msg), nil
}

func NewServer(ctx *internal.Context) *Server {
	return &Server{
		Ctx:  ctx,
		Repo: repository.NewMessageRepo(ctx),
	}
}
