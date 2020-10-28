package commands

import (
	"fmt"
	"github.com/rafaelbreno/go-bot/cmd/command"
	"github.com/rafaelbreno/go-bot/cmd/connection"
	"github.com/rafaelbreno/go-bot/cmd/helpers"
	"net"
)

func RunCommand(command string, conn net.Conn) {
	channel := connection.GetEnv("CHANNEL_NAME")
	GetInfo()
	switch command {
	case "!hello":
		helper.SendMessage("Hello, World", channel, conn)
		break
	default:
	}
}

func GetInfo() {
	commands := command.GetCommands()
	fmt.Println(commands)
}
