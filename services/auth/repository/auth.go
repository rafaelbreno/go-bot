package repository

import (
	"github.com/rafaelbreno/go-bot/auth/internal"
	"github.com/rafaelbreno/go-bot/auth/proto"
	"github.com/rafaelbreno/go-bot/auth/storage"
)

type AuthRepo interface {
	Create(req *proto.CreateRequest) *proto.CreateResponse
	Login(req *proto.LoginRequest) *proto.LoginResponse
	Check(req *proto.CheckRequest) *proto.CheckResponse
}

type AuthRepoCtx struct {
	Common  *internal.Common
	Storage *storage.Storage
}

func (a *AuthRepoCtx) Create(req *proto.CreateRequest) *proto.CreateResponse {
	return &proto.CreateResponse{}
}
func (a *AuthRepoCtx) Login(req *proto.LoginRequest) *proto.LoginResponse {
	return &proto.LoginResponse{}
}
func (a *AuthRepoCtx) Check(req *proto.CheckRequest) *proto.CheckResponse {
	return &proto.CheckResponse{}
}

func NewAuthRepo(c *internal.Common, sto *storage.Storage) *AuthRepoCtx {
	a := AuthRepoCtx{
		Common:  c,
		Storage: sto,
	}
	return &a
}
