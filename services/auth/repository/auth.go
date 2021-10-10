package repository

import (
	"fmt"

	"github.com/rafaelbreno/go-bot/auth/internal"
	"github.com/rafaelbreno/go-bot/auth/proto"
	"github.com/rafaelbreno/go-bot/auth/storage"
	"github.com/rafaelbreno/go-bot/auth/user"
	"golang.org/x/crypto/bcrypt"
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

const (
	fieldsNotMatch = `'%s' and '%s' don't match`
)

func (a *AuthRepoCtx) Create(req *proto.CreateRequest) *proto.CreateResponse {
	if req.Password != req.PasswordConfirmation {
		errMsg := fmt.Sprintf(fieldsNotMatch, "password", "password_confirmation")
		return &proto.CreateResponse{
			Token: "",
			Error: errMsg,
		}
	}

	encPW, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

	if err != nil {
		return &proto.CreateResponse{
			Error: err.Error(),
		}
	}

	u := user.User{
		Username: req.Username,
		Password: string(encPW),
	}

	if _, err := a.Storage.Pg.Conn.Model(u).Insert(); err != nil {
		return &proto.CreateResponse{
			Error: err.Error(),
		}
	}

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
