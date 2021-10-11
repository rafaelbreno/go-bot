package repository

import (
	"fmt"

	"github.com/go-pg/pg/v10"
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

func NewAuthRepo(c *internal.Common, sto *storage.Storage) *AuthRepoCtx {
	a := AuthRepoCtx{
		Common:  c,
		Storage: sto,
	}
	return &a
}

const (
	fieldsNotMatch = `'%s' and '%s' don't match`
	emptyField     = `field '%s' can not be empty`
	wrongPassord   = `wrong password`
)

func (a *AuthRepoCtx) Create(req *proto.CreateRequest) *proto.CreateResponse {
	if req.Username == "" {
		return &proto.CreateResponse{
			Error: fmt.Sprintf(emptyField, "username"),
		}
	}
	if req.Password == "" {
		return &proto.CreateResponse{
			Error: fmt.Sprintf(emptyField, "password"),
		}
	}
	if req.Password != req.PasswordConfirmation {
		errMsg := fmt.Sprintf(fieldsNotMatch, "password", "password_confirmation")
		return &proto.CreateResponse{
			Token: "",
			Error: errMsg,
		}
	}

	encPW, err := a.EncPassword(req.Password)

	if err != nil {
		return &proto.CreateResponse{
			Error: err.Error(),
		}
	}

	u := user.User{
		Username: req.Username,
		Password: string(encPW),
	}

	if _, err := a.Storage.Pg.Conn.Model(&u).Insert(); err != nil {
		return &proto.CreateResponse{
			Error: err.Error(),
		}
	}

	return &proto.CreateResponse{
		Token: "teste",
	}
}

func (a *AuthRepoCtx) Login(req *proto.LoginRequest) *proto.LoginResponse {
	return &proto.LoginResponse{}
}

func (a *AuthRepoCtx) Check(req *proto.CheckRequest) *proto.CheckResponse {
	return &proto.CheckResponse{}
}

func (a *AuthRepoCtx) EncPassword(pw string) (string, error) {
	encPW, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(encPW), err
}

func (a *AuthRepoCtx) CheckPassword(encPW, pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encPW), []byte(pw))
	return err == nil
}
