package main

import (
	"net"

	"github.com/rafaelbreno/go-bot/services/message-sender/internal"
	"github.com/rafaelbreno/go-bot/services/message-sender/proto"
	"github.com/rafaelbreno/go-bot/services/message-sender/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:5004")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	if err != nil {
		panic(err)
	}
	proto.RegisterSenderServer(grpcServer, server.NewServer(internal.NewContext()))

	grpcServer.Serve(lis)
}
