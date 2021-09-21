package server

import (
	"context"
	"time"

	"github.com/rafaelbreno/go-bot/services/message-reader/client"
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/proto"
	"google.golang.org/grpc"
)

type Server struct {
	Client *client.Client
	Ctx    *internal.Context
}

func NewServer(ctx *internal.Context) *Server {
	client := client.NewClient(ctx)

	return &Server{
		Client: client,
		Ctx:    ctx,
	}
}

func (s *Server) Start() {
	for {
		select {
		case msg := <-*s.Ctx.Msg:
			SendMessage(msg)
			//fmt.Println(msg.Channel)
			//fmt.Println(msg.Msg)
		}
	}
}

func SendMessage(msg *proto.MessageRequest) {
	conn, err := grpc.Dial("localhost:5004", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewSenderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)

	defer cancel()

	_, err = client.SendMessage(ctx, msg)

	if err != nil {
		panic(err)
	}
}

func (s *Server) Close() {
	s.Client.Conn.Close()
}
