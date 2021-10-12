package main

import (
	"net"

	"github.com/rafaelbreno/go-bot/auth/internal"
	"github.com/rafaelbreno/go-bot/auth/proto"
	"github.com/rafaelbreno/go-bot/auth/server"
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
	proto.RegisterAuthServer(grpcServer, server.NewServer(internal.NewCommon()))

	grpcServer.Serve(lis)
}
