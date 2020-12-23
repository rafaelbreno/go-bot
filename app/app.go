package app

import (
	"go-bot/app/command"
	"go-bot/cmd/env"
	"go-bot/cmd/irc"
)

/* App bootstrap struct
 * this will store the
 * main intern packages
**/
type App struct {
	Env env.Env
}

var app App

func init() {
	app = App{}
}

func Start() {
	conn := irc.GetConn()

	command.SetCommands(conn.Conn)

	defer conn.Disconnect()

	irc.Listen(conn.Conn)
}
