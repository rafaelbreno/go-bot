package app

import (
	"fmt"
	"go-bot/cmd/env"
	"go-bot/cmd/err"
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
	a := app.Env.Getenv("AABOT_USERNAME")

	fmt.Println(a)
}
