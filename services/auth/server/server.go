package server

import (
	"context"

	"github.com/rafaelbreno/go-bot/auth/internal"
	"github.com/rafaelbreno/go-bot/auth/proto"
	"github.com/rafaelbreno/go-bot/auth/repository"
	"github.com/rafaelbreno/go-bot/auth/storage"
)

type Server struct {
	proto.UnimplementedAuthServer
	Common *internal.Common
	Repo   *repository.AuthRepoCtx
}

func (s *Server) Create(ctx context.Context, req *proto.CreateRequest) *proto.CreateResponse {
	return s.Repo.Create(req)
}

func NewServer(common *internal.Common) *Server {
	return &Server{
		Common: common,
		Repo:   repository.NewAuthRepo(common, storage.NewStorage(common)),
	}
}
