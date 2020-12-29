package app

import (
	"go-bot/app/command"
	"go-bot/app/user"
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

	// Reading Commands JSON
	command.SetCommands(conn.Conn)

	// Importing
	user.ImportUsers()

	// When App is exited, disconnect from IRC
	defer conn.Disconnect()

	// When App is exited, save user data
	defer user.SaveUsers()

	// When App is exited, save user data
	defer command.SaveCommands()

	irc.Listen(conn.Conn)
}
