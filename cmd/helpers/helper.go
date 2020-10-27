package helper

import (
	"fmt"
	"net"
)

func SendMessage(message string, channel string, conn net.Conn) {
	conn.Write(buildMessage(message, channel))
}

func buildMessage(message string, channel string) []byte {
	msg := fmt.Sprintf("PRIVMSG #%s :%s", channel, message)
	return ParseToTwitch(msg)
}

func ParseToTwitch(str string) []byte {
	return []byte(fmt.Sprintf("%s\r\n", str))
}
