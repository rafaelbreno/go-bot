package command

import (
	"math/rand"
	"regexp"
	"time"

	"github.com/rafaelbreno/go-bot/internal"
)

// CommandCtx stores data to prepare
// the messages to be sent
type CommandCtx struct {
	Ctx *internal.Context
}

// Command store commands
type Command struct {
	Key     string
	Answer  string
	Options []string
	Type    int
}

type Action struct {
	SentBy string
}

var commands = map[string]Command{
	"!hello": {
		Key:    "!hello",
		Type:   Simple,
		Answer: "Hello, {user}!",
	},
	"!signo": {
		Key:     "!signo",
		Type:    Random,
		Answer:  "/me {user} decidiu trocar de signo, agora seu novo signo é: {answer}",
		Options: []string{"batata", "cadeira de massagem ", "brownie de feijão", "chuveiro frio", "carioca"},
	},
	"!cupido": {
		Key:    "!cupido",
		Type:   Cupido,
		Answer: "/me {user} sua alma gêmea é: @{user_list}",
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
	rand.Seed(time.Now().Unix())

	switch c.Type {
	case Simple:
		return replace(c.Answer, keyMap{
			"{user}": act.SentBy,
		})
	case Random:
		if act.SentBy == "rafiusky" {
			return replace(c.Answer, keyMap{
				"{user}":   act.SentBy,
				"{answer}": "O Glorioso",
			})
		}
		return replace(c.Answer, keyMap{
			"{user}":   act.SentBy,
			"{answer}": random(c.Options),
		})
	case Cupido:
		ans := ""
		if val, ok := cupidPair[act.SentBy]; ok {
			ans = val
		} else {
			ans = random(H.fetchUserList(), append(modBlacklist, "lajurubeba", "rafiusky", "rafiuskybot", act.SentBy)...)
		}
		return replace(c.Answer, keyMap{
			"{user}":      act.SentBy,
			"{user_list}": ans,
		})
	default:
		return ""
	}
}
