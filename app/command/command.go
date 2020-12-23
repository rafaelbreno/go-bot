package command

import (
	"fmt"
	"go-bot/cmd/helper"
	"net"
	"strings"
)

type Message struct {
	Command  string
	Channel  string
	Username string
	Message  string
	Response Response
	Twitch   bool
}

type Response struct {
	Body string
}

type Command struct {
	Trigger  string
	Response string
}

var commands map[string]Command
var conn net.Conn

func SetCommands(c net.Conn) {
	conn = c

	commands = make(map[string]Command, 10)

	commands["!hello"] = Command{
		Trigger:  "!hello",
		Response: "Hello __USER__",
	}
}

func ChatCommand(m Message) {
	c := strings.SplitN(m.Message, " ", 1)

	fmt.Println("Command found:", c[0])

	if _, found := commands[c[0]]; found {
		m.sendCommand(commands[c[0]].prepareResponse(m))
	}
}

func (m Message) sendCommand(r string) {
	str := fmt.Sprintf("PRIVMSG #%s :%s", m.Channel, r)

	helper.WriteTwitch(str, conn)
}

func (c Command) prepareResponse(m Message) string {
	resp := c.Response

	resp = strings.ReplaceAll(resp, "__USER__", "@"+m.Username)

	return resp
}
