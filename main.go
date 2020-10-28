package main

import (
	"bufio"
	"fmt"
	"github.com/rafaelbreno/go-bot/cmd/actions"
	"github.com/rafaelbreno/go-bot/cmd/connection"
	"github.com/rafaelbreno/go-bot/cmd/helpers"
	"github.com/rafaelbreno/go-bot/cmd/parser"
	"strings"

	"net"
	"net/textproto"
)

func sendMessage(channel string, message string) []byte {
	msg := fmt.Sprintf("PRIVMSG #%s :%s\r\n", channel, message)
	return helper.ParseToTwitch(msg)
}

func execMessage(conn net.Conn, message parser.Message) {
	fmt.Println(message)

	if strings.HasPrefix(message.Command, "!") {
		commands.RunCommand(message.Command, conn)
	}
	if message.Command == "PONG" {
		conn.Write(helper.ParseToTwitch(message.Content))
	}
}

func listen() {
	conn := connection.Connect()
	connection.Logon(conn)
	tp := textproto.NewReader(bufio.NewReader(conn))
	for {
		status, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		message := parser.Parse(status)
		execMessage(conn, message)
	}
	connection.Disconnect(conn)
}

func main() {
	listen()
}
