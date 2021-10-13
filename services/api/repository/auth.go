package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/rafaelbreno/go-bot/api/entity"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/proto"
	grpc "google.golang.org/grpc"
)

type AuthRepo interface {
	Create(u entity.User) (entity.UserResponse, error)
	Login(u entity.User) (entity.UserResponse, error)
	Check(token string) error
}

type AuthRepoCtx struct {
	Ctx        *internal.Context
	AuthClient proto.AuthClient
}

func NewAuthRepoCtx(ctx *internal.Context) *AuthRepoCtx {
	conn, err := grpc.Dial("localhost:5004", grpc.WithInsecure())

	if err != nil {
		ctx.Logger.Error(err.Error())
		return &AuthRepoCtx{}
	}

	client := proto.NewAuthClient(conn)

	return &AuthRepoCtx{
		Ctx:        ctx,
		AuthClient: client,
	}
}

func (a *AuthRepoCtx) Create(u entity.User) (entity.UserResponse, error) {
	createReq := &proto.CreateRequest{
		Username:             u.Username,
		Password:             u.Password,
		PasswordConfirmation: u.PasswordConfirmation,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()

	createRes, err := a.AuthClient.Create(ctx, createReq)

	if err != nil {
		return entity.UserResponse{
			Error: err.Error(),
		}, err
	}

	userResp := entity.UserResponse{
		Token: createRes.Token,
		Error: createRes.Error,
	}

	return userResp, nil
}

func (a *AuthRepoCtx) Login(u entity.User) (entity.UserResponse, error) {
	loginReq := &proto.LoginRequest{
		Username: u.Username,
		Password: u.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()

	loginRes, err := a.AuthClient.Login(ctx, loginReq)

	if err != nil {
		return entity.UserResponse{
			Error: err.Error(),
		}, err
	}

	userResp := entity.UserResponse{
		Token: loginRes.Token,
		Error: loginRes.Error,
	}

	return userResp, nil
}

func (a *AuthRepoCtx) Check(token string) error {
	checkReq := &proto.CheckRequest{
		Token: token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()

	checkRes, err := a.AuthClient.Check(ctx, checkReq)

	if err != nil {
		return err
	}

	if checkRes.Error != "" {
		return fmt.Errorf(checkRes.Error)
	}

	return nil
}
