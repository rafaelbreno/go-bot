package command

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
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
		Answer:  "Seu novo signo {answer}",
		Options: []string{"Capricórnio", "Sagitário", "Escorpião", "Leão", "batata", "cadeira de massagem", "brownie de feijão", "chuveiro frio", "carioca"},
	},
	"!cupido": {
		Key:    "!cupido",
		Type:   Cupido,
		Answer: "Sua alma gêmea é: @{user_list}",
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
				"{answer}": "O Glorioso",
			})
		}
		return replace(c.Answer, keyMap{
			"{answer}": random(c.Options),
		})
	case Cupido:
		ans := ""
		switch act.SentBy {
		case "lajurubeba":
			ans = "rafiusky"
		case "rafiusky":
			ans = "lajurubeba"
		default:
			ans = random(fetchUserList())
		}
		return replace(c.Answer, keyMap{
			"{user_list}": ans,
		})
	default:
		return ""
	}
}

type TMIViewers struct {
	Chatters struct {
		Broadcaster []string `json:"broadcaster"`
		Vips        []string `json:"vips"`
		Moderators  []string `json:"moderators"`
		Viewers     []string `json:"viewers"`
	} `json:"chatters"`
}

func fetchUserList() []string {
	v := TMIViewers{}
	resp, _ := http.Get("https://tmi.twitch.tv/group/user/rafiusky/chatters")
	_ = json.NewDecoder(resp.Body).Decode(&v)

	var list []string
	list = append(list, v.Chatters.Broadcaster...)
	list = append(list, v.Chatters.Viewers...)
	list = append(list, v.Chatters.Vips...)
	list = append(list, v.Chatters.Moderators...)
	return list
}

func random(list []string) string {
	rand.Seed(time.Now().Unix())
	return list[rand.Intn(len(list))]
}

func replace(str string, repMap keyMap) string {
	for key, val := range repMap {
		str = strings.ReplaceAll(str, key, val)
	}
	return str
}
