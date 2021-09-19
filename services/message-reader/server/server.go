package server

import (
	"fmt"

	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/proto"
	"google.golang.org/grpc"
)

type Server struct {
	Msg    chan *proto.MessageRequest
	Client proto.SenderClient
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
	}
}

func (s *Server) Start() {
	for {
		select {
		case msg := <-s.Msg:
			fmt.Println(msg)
		}
	}
}
