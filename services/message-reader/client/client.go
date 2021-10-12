package client

import (
	"github.com/rafaelbreno/go-bot/services/message-reader/internal"
	"github.com/rafaelbreno/go-bot/services/message-reader/proto"
	"google.golang.org/grpc"
)

type Client struct {
	Ctx        *internal.Context
	Conn       *grpc.ClientConn
	GRPCClient proto.SenderClient
}

func NewClient(ctx *internal.Context) *Client {
	service_url := string(ctx.Env["SENDER_SERVICE_URL"] + ":" + ctx.Env["SENDER_SERVICE_PORT"])

	conn, err := grpc.Dial(service_url, grpc.WithInsecure())

	client := proto.NewSenderClient(conn)

	if err != nil {
		ctx.Logger.Fatal(err.Error())
	}

	return &Client{
		Ctx:        ctx,
		GRPCClient: client,
		Conn:       conn,
	}
}
