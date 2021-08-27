package utils

import (
	"fmt"
	"net"

	"github.com/rafaelbreno/go-bot/internal"
)

func Write(ctx *internal.Context, conn net.Conn, msg string) {
	ctx.Logger.Info("Sending Message")
	_, err := conn.Write([]byte(fmt.Sprintf("%s\r\n", msg)))
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
}
