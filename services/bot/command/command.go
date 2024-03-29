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
	H   *CmdHelper
}

// Command store commands
type Command struct {
	Key         string
	Answer      string
	Options     []string
	Type        int
	HasCooldown bool
	Cooldown    time.Duration
	ExpireAt    int64
}

type Action struct {
	SentBy string
}

var commands = map[string]*Command{
	"!hello": {
		Key:         "!hello",
		Type:        Simple,
		Answer:      "Hello, {user}!",
		HasCooldown: true,
		Cooldown:    time.Duration(15 * time.Second),
		ExpireAt:    0,
	},
	"!signo": {
		Key:         "!signo",
		Type:        Random,
		HasCooldown: false,
		Answer:      "/me {user} decidiu trocar de signo, agora seu novo signo é: {answer}",
		Options:     lstSigno,
	},
	"!cupido": {
		Key:         "!cupido",
		Type:        Cupido,
		HasCooldown: false,
		Answer:      "/me {user} sua alma gêmea é: @{user_list}",
	},
}

var (
	cmdRegex = regexp.MustCompile(`^(!)[a-zA-Z0-9]{1,}`)
)

// GetAnswer receives a message to be sent
func (c *CommandCtx) GetAnswer(sentBy, inMessage string) string {
	cmdKey := string(cmdRegex.Find([]byte(inMessage)))

	var cmd *Command
	var ok bool

	if cmd, ok = commands[cmdKey]; !ok {
		return ""
	}

	action := &Action{
		SentBy: sentBy,
	}

	return cmd.prepare(action, c)
}

type keyMap map[string]string

func (c *Command) prepare(act *Action, ctx *CommandCtx) string {
	timeNow := time.Now()
	timeNowUnix := timeNow.Unix()

	rand.Seed(timeNowUnix)

	if c.HasCooldown {
		if !(c.ExpireAt == 0 || c.ExpireAt <= timeNowUnix) {
			return ""
		}
		c.ExpireAt = timeNow.Add(c.Cooldown).Unix()
	}

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
		if act.SentBy == "lajurubeba" {
			return replace(c.Answer, keyMap{
				"{user}":   act.SentBy,
				"{answer}": "espirro de loli",
			})
		}
		return replace(c.Answer, keyMap{
			"{user}":   act.SentBy,
			"{answer}": random(c.Options),
		})
	case Cupido:
		ans := ""

		ans = random(ctx.H.fetchUserList(), append(modBlacklist, "lajurubeba", "rafiusky", "rafiuskybot", act.SentBy)...)

		return replace(c.Answer, keyMap{
			"{user}":      act.SentBy,
			"{user_list}": ans,
		})
	default:
		return ""
	}
}
