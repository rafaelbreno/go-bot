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

func (s *Server) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	return s.Repo.Create(req), nil
}

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	return s.Repo.Login(req), nil
}

func (s *Server) Check(ctx context.Context, req *proto.CheckRequest) (*proto.CheckResponse, error) {
	return s.Repo.Check(req), nil
}

func NewServer(common *internal.Common) *Server {
	return &Server{
		Common: common,
		Repo:   repository.NewAuthRepo(common, storage.NewStorage(common)),
	}
}
