package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rafaelbreno/go-bot/cmd/connection"
	"github.com/rafaelbreno/go-bot/parser"

	"net"
	"net/textproto"
	"os"
)

func sendMessage(channel string, message string) []byte {
	msg := fmt.Sprintf("PRIVMSG #%s :%s\r\n", channel, message)
	return []byte(msg)
}

func sendPong(command string) []byte {
	return []byte(command)
}

func execMessage(conn net.Conn, message parser.Message) {
	if message.Command == "PONG" {
		conn.Write([]byte(message.Content))
	}
}

func main() {
	conn := connection.Connect()
	logon(conn)
	tp := textproto.NewReader(bufio.NewReader(conn))
	for {
		status, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Println(status)
		message := parser.Parse(status)
		execMessage(conn, message)
		fmt.Println(message)
		//		fmt.Printf("Type: %T \n Status: %s \n\n", status, status)
	}
	disconnect(conn)
}
