package server

import (
	"context"
	"fmt"

	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/proto"
	"google.golang.org/grpc"
)

type Server struct {
	Msg    chan *proto.MessageRequest
	Client proto.SenderClient
	Ctx    *internal.Context
}

func NewServer(ctx *internal.Context) *Server {
	addr := "localhost:5004"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		ctx.Logger.Panic(err.Error())
	}

	client := proto.NewSenderClient(conn)

	return &Server{
		Msg:    make(chan *proto.MessageRequest),
		Client: client,
		Ctx:    ctx,
	}
}

func (s *Server) Start() {
	for {
		select {
		case msg := <-s.Msg:
			if _, err := s.Client.SendMessage(context.Background(), msg, nil); err != nil {
				s.Ctx.Logger.Panic(err.Error())
			}
			fmt.Println(msg.Channel)
			fmt.Println(msg.Msg)
		}
	}
}
