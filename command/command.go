package command

import (
	"regexp"
	"strings"

	"github.com/rafaelbreno/go-bot/internal"
)

// CommandCtx stores data to prepare
// the messages to be sent
type CommandCtx struct {
	Ctx *internal.Context
}

// Command store commands
type Command struct {
	Key    string
	Answer string
}

type Action struct {
	SentBy string
}

var commands = map[string]Command{
	"!hello": Command{
		Key:    "!hello",
		Answer: "Hello, {user}!",
	},
}

var (
	cmdRegex = regexp.MustCompile(`^(!)[a-zA-Z0-9]{1,}`)
)

// GetAnswer receives a message to be sent
func (c *CommandCtx) GetAnswer(sentBy, inMessage string) string {
	cmdKey := string(cmdRegex.Find([]byte(inMessage)))

	var cmd Command
	var ok bool

	if cmd, ok = commands[cmdKey]; !ok {
		return ""
	}

	action := &Action{
		SentBy: sentBy,
	}

	return cmd.prepare(action)
}

type keyMap map[string]string

func (c *Command) prepare(act *Action) string {
	return replace(c.Answer, keyMap{
		"{user}": act.SentBy,
	})
}

func replace(str string, repMap keyMap) string {
	for key, val := range repMap {
		str = strings.ReplaceAll(str, key, val)
	}
	return str
}
