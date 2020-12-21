package app

import (
	"go-bot/cmd/env"
	"go-bot/cmd/err"
	"go-bot/cmd/irc"
)

/* App bootstrap struct
 * this will store the
 * main intern packages
**/
type App struct {
	Env env.Env
	Err *err.Err
}

var app App

func init() {
	app = App{
		Err: &err.Err{},
	}
	app.Env.Err = app.Err
}

func Start() {
	conn := irc.GetConn()

	defer conn.Disconnect()
}
