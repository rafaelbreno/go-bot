package main

import (
	"context"
	"fmt"
	"grpc-client/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5004", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewSenderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)

	defer cancel()

	resp, err := client.SendMessage(ctx, &proto.MessageRequest{
		Msg:     "salve fiote",
		Channel: "rafiusky",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
