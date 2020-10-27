package commands

import (
	"github.com/rafaelbreno/go-bot/cmd/connection"
	"github.com/rafaelbreno/go-bot/cmd/helpers"
	"net"
)

func Command(command string, conn net.Conn) {
	channel := connection.GetEnv("CHANNEL_NAME")
	switch command {
	case "!hello":
		helper.SendMessage("Hello, World", channel, conn)
		break
	default:
	}
}
