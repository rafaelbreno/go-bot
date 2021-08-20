package command

import (
	"github.com/rafaelbreno/go-bot/internal"
)

type Command struct {
}

var ctx *internal.Context

// Start a
func Start(contx *internal.Context) {
	ctx = contx
}
