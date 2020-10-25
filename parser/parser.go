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

func parseUserMessage(rawStatus string) Message {
	message := Message{}
	splitted := strings.SplitN(rawStatus, "!", 2)
	// Getting Username
	message.Username = strings.TrimPrefix(splitted[0], ":")
	message.Content = splitted[1]
	return message
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
	if strings.HasPrefix(rawStatus, ":tmi.twitch.tv") || strings.HasPrefix(rawStatus, ":rafiuskybot.tmi.twitch.tv") {
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
		message = parseUserMessage(rawStatus)
	}
	return message
}

func Parse(rawStatus string) Message {
	return split(rawStatus)
}
