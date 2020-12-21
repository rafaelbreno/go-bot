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
	Err err.Err
}

var app App

func Start() {
	a, err := app.Env.Getenv("BOT_USERNAME")
	if err != nil {
		app.Err.Log(err)
	}

	fmt.Println(a)
}
