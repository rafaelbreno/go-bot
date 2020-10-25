package parser

import (
	//	"fmt"
	"strings"
)

type Message struct {
	Username string
	Content  string
	Command  string
}

/*
Possible Messages:
- Twitch default connecting messages
	:tmi.twitch.tv 376 rafiuskybot
- User message
	:rafiusky!rafiusky@rafiusky.tmi.twitch.tv PRIVMSG #rafiusky :a
- Twitch PING
	PING :tmi.twitch.tv

*/
func split(rawStatus string) Message {

	message := Message{}
	if strings.HasPrefix(rawStatus, ":tmi.twitch.tv") {
		// Default twitch message
		message.Username = "Twitch"
		message.Content = ""
		message.Command = ""
	} else if strings.HasPrefix(rawStatus, "PING") {
		// Twitch Ping
		message.Username = "Twitch"
		message.Content = "PONG :tmi.twitch.tv"
		message.Command = "PONG"
	} else {
		// User's Message
		message.Username = "Twitch"
		message.Content = "User Message"
		message.Command = ""
	}
	return message
}

func Parse(rawStatus string) Message {
	return split(rawStatus)
}
