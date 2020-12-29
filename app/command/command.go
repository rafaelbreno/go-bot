package command

import (
	"encoding/json"
	"fmt"
	"go-bot/app/user"
	"go-bot/cmd/helper"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

type Message struct {
	Command  string
	Channel  string
	Username string
	Message  string
	Response Response
	Twitch   bool
}

type Response struct {
	Body string
}

type Commands struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	Trigger  string `json:"trigger"`
	Response string `json:"response"`
}

var commands Commands
var conn net.Conn

func SetCommands(c net.Conn) {
	conn = c

	jsonFile, err := os.Open("db/commands.json")

	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(jsonBytes, &commands); err != nil {
		log.Fatal(err)
	}
}

// Refactor
func findCommand(comm string) (Command, bool) {
	for _, value := range commands.Commands {
		if value.Trigger == comm {
			return value, true
		}
	}

	return Command{}, false
}

func ChatCommand(m Message) {
	user.GetUser(m.Username)

	c := strings.SplitN(m.Message, " ", 1)

	fmt.Println("Command found:", c[0])

	if comm, found := findCommand(c[0]); found {
		m.sendCommand(comm.prepareResponse(m))
	}
}

func (m Message) sendCommand(r string) {
	str := fmt.Sprintf("PRIVMSG #%s :%s", m.Channel, r)

	helper.WriteTwitch(str, conn)
}

func (c Command) prepareResponse(m Message) string {
	resp := c.Response

	resp = strings.ReplaceAll(resp, "__USER__", "@"+m.Username)

	return resp
}

func SaveCommands() {
	log.Println("Commands Saved!")
}
