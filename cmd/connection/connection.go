package connection

import (
	"fmt"
	"github.com/joho/godotenv"
	"net"
	"os"
)

type Bot struct {
	ConnectionURL  string
	ConnectionPort string

	BotUsername string
	BotPassword string
	BotToken    string

	Channel string
}

func Connect() net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		panic(err)
	}
	return conn
}

func disconnect(conn net.Conn) {
	conn.Close()
}

func logon(conn net.Conn) {
	bot := getCredentials()

	pass := []byte(fmt.Sprintf("PASS %s\r\n", bot.BotToken))
	nick := []byte(fmt.Sprintf("NICK %s\r\n", bot.BotUsername))
	chann := []byte(fmt.Sprintf("JOIN #%s\r\n", bot.Channel))

	conn.Write(pass)
	conn.Write(nick)
	conn.Write(chann)
}

func getCredentials() Bot {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	bot := Bot{}
	bot.ConnectionURL = "irc.chat.twitch.tv"
	bot.ConnectionPort = "6667"
	bot.BotUsername = os.Getenv("BOT_USERNAME")
	bot.BotToken = os.Getenv("OAUTH_TOKEN")
	bot.Channel = os.Getenv("CHANNEL_NAME")

	return bot
}
