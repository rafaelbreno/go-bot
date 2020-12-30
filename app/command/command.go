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
	"regexp"
	"strings"
)

const (
	COMMAND_VAR_REGEXP = `` +
		// . Any character
		// ? Any lenght
		`.?` +

		// Need to start with '__'
		`__` +

		/* Need to have at least one of:
		 * a-z Character
		 * A-Z Character
		 * @ Character
		 * ( Character
		 * ) Character
		**/
		`[a-zA-z@\(\)]{1,}` +

		// Need to end with '__'
		`__`
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

type Parser struct {
	// __Sender@Username__ has __Sender@Coins(int)__ coins
	RawString string

	// [__Sender@Username__ __Sender@Coins(int)__ ]
	Keys []Key

	// @rafiusky has 200 coins
	PreparedString string

	MessageData Message
}

type Key struct {
	// __Sender@Username__
	RawKey string

	// rafiusky
	PreparedKey string
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
	p := Parser{
		MessageData: m,
		RawString:   c.Response,
	}

	p.Parse()

	log.Println("prepareResponse", p.PreparedString)

	//resp := c.Response

	//resp = strings.ReplaceAll(resp, "__USER__", "@"+m.Username)

	return p.PreparedString
}

func (p *Parser) Parse() {
	p.
		setKeys().
		prepareKeys()
}

func (p *Parser) setKeys() *Parser {
	re := regexp.MustCompile(COMMAND_VAR_REGEXP)
	foundKeys := re.FindAllString(p.RawString, -1)

	for _, value := range foundKeys {
		p.Keys = append(p.Keys, Key{
			RawKey: value,
		})
	}

	return p
}

func (p *Parser) prepareKeys() *Parser {
	p.PreparedString = p.RawString

	for _, value := range p.Keys {
		var err error
		// Sender@Username
		trimmed := strings.Trim(value.RawKey, " ")
		trimmed = strings.Trim(trimmed, "__")

		// [Sender Username]
		parts := strings.Split(trimmed, "@")

		value.PreparedKey, err = p.getValue(parts)

		if err != nil {
			panic(err)
		}

		p.PreparedString = strings.ReplaceAll(p.PreparedString, value.RawKey, value.PreparedKey)
	}
	return p
}

func (p Parser) getValue(parts []string) (string, error) {
	switch parts[0] {
	// Who sent the message
	case "Target":
		fallthrough
	case "Sender":
		// User struct tags
		return user.GetByTag(p.MessageData.Username, parts[1])
	default:
		return "", fmt.Errorf("Key %s not found!", parts[0])
	}
}

func SaveCommands() {
	log.Println("Commands Saved!")
}
