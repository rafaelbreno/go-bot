package main

import (
	"bufio"
	"fmt"
	"github.com/rafaelbreno/go-bot/cmd/connection"
	"github.com/rafaelbreno/go-bot/parser"

	"net"
	"net/textproto"
)

func parseToTwitch(str string) []byte {
	return []byte(fmt.Sprintf("%s\r\n", str))
}

func sendMessage(channel string, message string) []byte {
	msg := fmt.Sprintf("PRIVMSG #%s :%s", channel, message)
	return parseToTwitch(msg)
}

func execMessage(conn net.Conn, message parser.Message) {
	fmt.Println(message)
	//	fmt.Println(message.Username == "Twitch")
	if message.Command == "PONG" {
		conn.Write(parseToTwitch(message.Content))
	}
}

func main() {
	conn := connection.Connect()
	connection.Logon(conn)
	tp := textproto.NewReader(bufio.NewReader(conn))
	for {
		status, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		//	fmt.Println(status)
		message := parser.Parse(status)
		execMessage(conn, message)
		//	fmt.Println(message)
	}
	connection.Disconnect(conn)
}
