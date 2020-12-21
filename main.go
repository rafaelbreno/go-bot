package main

import (
	"go-bot/app"
	"go-bot/cmd/helper"
)

func main() {
	app.Start()

	helper.WaitForCtrlC()
}
