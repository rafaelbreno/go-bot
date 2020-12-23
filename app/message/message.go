package message

import (
	"fmt"
	"go-bot/app/command"
	"net"
	"regexp"
	"strings"
)

const (
	// Key to dump any value
	TEMPKEY = "temp"

	/* Must start with :
	 * Can have N charactes, digits and underscore
	 * Must have one !
	 * Username
	**/
	USERNAMEKEY    = "username"
	USERNAME1REGEX = `(:)[\w+]{1,}(!)[\w+]{1,}(@)[\w+]{1,}(.tmi.twitch.tv)`

	/* Can have N charactes, digits and underscore, followed by '#'
	 * Command
	**/
	COMMANDKEY   = "command"
	COMMANDREGEX = `[\w\s]{1,}(#)`

	/* Can have N charactes, digits and underscore, followed by ':'
	 * Username
	**/
	CHANNELKEY   = "channel"
	CHANNELREGEX = `[\w\s]{1,}(:)`

	// User message
	MESSAGEKEY   = "message"
	MESSAGEREGEX = `[\w\s\S]{1,}`

	IRCREGEX = `(:)[\w+]{1,}(!)[\w+]{1,}(@)[\w+]{1,}(.tmi.twitch.tv )[\w]{1,}( #)[\w]{1,}( :)[\s\S]{1,}`
)

type Message struct {
	Raw      string
	TempRaw  string
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

type Steps struct {
	Regex string
	Key   string
}

var steps []Steps
var conn net.Conn

func SetSteps(conn net.Conn) {
	steps = make([]Steps, 4)

	steps[0] = Steps{
		Regex: USERNAME1REGEX,
		Key:   USERNAMEKEY,
	}

	steps[1] = Steps{
		Regex: COMMANDREGEX,
		Key:   COMMANDKEY,
	}

	steps[2] = Steps{
		Regex: CHANNELREGEX,
		Key:   CHANNELKEY,
	}

	steps[3] = Steps{
		Regex: MESSAGEREGEX,
		Key:   MESSAGEKEY,
	}
}

func ReadAndParse(msg string) {
	m := Message{
		Raw:     msg,
		TempRaw: msg,
	}

	if err := m.Parse(); err != nil {
		panic(err)
	}

	fmt.Printf("%s disse: %s\n\n", m.Username, m.Message)

	if strings.HasPrefix(m.Message, "!") {
		comm := command.Message{
			Channel:  m.Channel,
			Command:  m.Command,
			Username: m.Username,
			Message:  m.Message,
			Response: command.Response{
				Body: m.Response.Body,
			},
			Twitch: m.Twitch,
		}

		go command.ChatCommand(comm)
	}
}

/*
 * User Message
	:rafiusky!rafiusky@rafiusky.tmi.twitch.tv PRIVMSG #rafiusky :Hello World
	:rafiuskybot!rafiuskybot@rafiuskybot.tmi.twitch.tv PRIVMSG #rafiusky :A
	:<user>!<user>@<user>.tmi.twitch.tv <command> #channel :<msg>

 * Ping
 	PING :tmi.twitch.tv
**/

func (m *Message) Parse() error {
	if ok := regexp.MustCompile(`(PING)`).Match([]byte(m.Raw)); ok {
		m.Twitch = true
		// Message to continue connected to Twitch's IRC
		m.Message = "PONG"
		return nil
	}

	if ok, err := regexp.Match(IRCREGEX, []byte(m.Raw)); !ok || err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	for _, value := range steps {

		str := regexp.MustCompile(value.Regex).FindString(m.TempRaw)

		switch value.Key {
		case "username":
			/*
				:<username>!
			*/
			m.TempRaw = strings.TrimPrefix(m.TempRaw, str)

			usernameRaw := strings.Split(str, "!")
			username := strings.Trim(usernameRaw[0], ":")
			m.Username = username

			break

		case "command":
			/*
				<COMMAND>
			*/
			m.TempRaw = strings.TrimPrefix(m.TempRaw, str)

			command := strings.TrimPrefix(str, " ")
			command = strings.TrimSuffix(str, " #")
			m.Command = command

			break

		case "channel":
			/*
				#<CHANNEL>
			*/
			m.TempRaw = strings.TrimPrefix(m.TempRaw, str)

			channel := strings.Trim(str, ":")
			channel = strings.Trim(channel, " ")
			m.Channel = channel
			break

		case "message":
			m.TempRaw = strings.TrimPrefix(m.TempRaw, str)

			m.Message = str
			break

		default:
			break
		}
	}
	return nil
}
