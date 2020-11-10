package commands

import (
	"fmt"
	"github.com/rafaelbreno/go-bot/cmd/command"
	"github.com/rafaelbreno/go-bot/cmd/connection"
	"github.com/rafaelbreno/go-bot/cmd/helpers"
	"net"
	"sort"
)

var channel string
var commands []command.Command

func RunCommand(command string, conn net.Conn) {
	channel = connection.GetEnv("CHANNEL_NAME")

	commands = GetInfo()

	fmt.Println("Commands: ", commands)

	i := sort.Search(len(commands), func(i int) bool {
		return command == commands[i].Identifier
	})

	fmt.Println(i)

}

func ExecCommand(channel string, conn net.Conn) {
	helper.SendMessage("Hello, World", channel, conn)
}

func GetInfo() []command.Command {
	return command.GetCommands()
}
