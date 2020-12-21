package irc

import (
	"fmt"
	"go-bot/cmd/app_error"
	"go-bot/cmd/env"
	"net"
)

type Conn struct {
	Cred Credentials
	Conn net.Conn
	Env  env.Env
}

type Credentials struct {
	ConnURL      string
	ConnPort     string
	ConnProtocol string

	BotUsername   string
	BotPassword   string
	BotOAuthToken string

	TwitchChannel string
}

func GetConn() *Conn {
	conn := Conn{
		Env: env.Env{},
	}

	conn.
		setCredentials().
		setConn().
		connectIRC()

	return &conn
}

func (c *Conn) setCredentials() *Conn {
	cred := Credentials{
		ConnURL:       "irc.chat.twitch.tv",
		ConnPort:      "6667",
		ConnProtocol:  "tcp",
		BotUsername:   c.Env.Getenv("BOT_USERNAME"),
		BotOAuthToken: c.Env.Getenv("OAUTH_TOKEN"),
		TwitchChannel: c.Env.Getenv("CHANNEL_NAME"),
	}
	c.Cred = cred

	return c
}

//
func (c *Conn) setConn() *Conn {
	conn, e := net.
		Dial(c.Cred.ConnProtocol,
			fmt.Sprintf("%s:%s", c.Cred.ConnURL, c.Cred.ConnPort))

	if e != nil {
		app_error.NewError(e.Error(), "cmd/irc/irc.go:setConn")
	}

	c.Conn = conn

	return c
}

/* Setting connection with Channel's Chat IRC */
func (c *Conn) connectIRC() {
	c.Write(fmt.Sprintf("PASS %s", c.Cred.BotOAuthToken))
	c.Write(fmt.Sprintf("NICK %s", c.Cred.BotUsername))
	c.Write(fmt.Sprintf("JOIN #%s", c.Cred.TwitchChannel))
}

func (c *Conn) Disconnect() {
	if e := c.Conn.Close(); e != nil {
		app_error.NewError(e.Error(), "cmd/irc/irc.go:Disconnect")
	}
}

/* Helper to write IRC messages
 * following the default structure:
 * MSG \r\n
**/
func (c *Conn) Write(msg string) {
	c.Conn.Write([]byte(fmt.Sprintf("%s\r\n", msg)))
}
