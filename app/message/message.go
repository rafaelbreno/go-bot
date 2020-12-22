package message

import (
	"fmt"
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
	USERNAME1REGEX = `(:)[\w+]{1,}(!)`

	/* Can have N charactes, digits and underscore
	 * Followed by '@'
	 * Username
	**/

	USERNAME2REGEX = `[\w+]{1,}(@)`

	/* Can have N charactes, digits and underscore
	 * Followed by '.tmi.twitch.tv'
	 * Username
	**/
	USERNAME3REGEX = `[\w+]{1,}(.tmi.twitch.tv)`

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
	MESSAGEREGEX = `[\s\S]{1,}`

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

var steps map[int]Steps

func SetSteps() {
	steps = make(map[int]Steps, 6)

	steps[0] = Steps{
		Regex: USERNAME1REGEX,
		Key:   USERNAMEKEY,
	}

	steps[1] = Steps{
		Regex: USERNAME2REGEX,
		Key:   TEMPKEY,
	}

	steps[2] = Steps{
		Regex: USERNAME3REGEX,
		Key:   TEMPKEY,
	}

	steps[3] = Steps{
		Regex: COMMANDREGEX,
		Key:   COMMANDKEY,
	}

	steps[4] = Steps{
		Regex: CHANNELREGEX,
		Key:   CHANNELKEY,
	}

	steps[5] = Steps{
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

	//fmt.Println(m.Raw)

	fmt.Println("Channel: ", m.Channel)
	fmt.Println("Message: ", m.Message)
	fmt.Println("Command: ", m.Command)
	fmt.Println("Username: ", m.Username)
}

/*
 * User Message
	:rafiusky!rafiusky@rafiusky.tmi.twitch.tv PRIVMSG #rafiusky :Hello World
	:rafiuskybot!rafiuskybot@rafiuskybot.tmi.twitch.tv PRIVMSG #rafiusky :A
	:<user>!<user>@<user>.tmi.twitch.tv <command> #channel :<msg>

 * Ping
 	PING :tmi.twitch.tv
**/

func (m Message) Parse() error {
	//fmt.Println("AAAAAAAAAAAAAaaa")

	if ok := regexp.MustCompile(`(PING)`).Match([]byte(m.Raw)); ok {
		//fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
		m.Twitch = true
		// Message to continue connected to Twitch's IRC
		m.Message = "PONG"
		return nil
	}

	if ok, err := regexp.Match(IRCREGEX, []byte(m.Raw)); !ok || err != nil {
		//fmt.Println("CCCCCCCCCCCCCCCCCCCCCCCCCCC")
		fmt.Println("Error: ", err)
		return err
	}

	for _, value := range steps {
		//fmt.Println("DDDDDDDDDDDDDDDDDDDDDDDDD")

		str := regexp.MustCompile(value.Regex).FindString(m.TempRaw)

		fmt.Println(value.Key)

		switch value.Key {
		case "username":
			/*
				:<username>!
			*/
			//fmt.Println("Username", str)
			username := strings.Trim(str, ":")
			username = strings.Trim(username, "!")
			m.Username = username

			m.TempRaw = strings.Trim(m.TempRaw, str)

			break

		case "command":
			/*
				<COMMAND>
			*/
			//fmt.Println("Command", str)

			command := strings.Trim(str, " ")
			m.Command = command

			m.TempRaw = strings.Trim(m.TempRaw, str)
			break

		case "channel":
			/*
				#<CHANNEL>
			*/
			//fmt.Println("Channel", str)

			channel := strings.Trim(str, "#")
			channel = strings.Trim(channel, " ")
			m.Channel = channel

			m.TempRaw = strings.Trim(m.TempRaw, str)

			break

		case "message":
			//fmt.Println("Message", str)

			msg := strings.TrimPrefix(str, ":")
			m.Message = msg

			m.TempRaw = strings.Trim(m.TempRaw, str)
			break

		default:
			/*
			 *	random
			**/
			//fmt.Println("Random", str)
			m.TempRaw = strings.Trim(m.TempRaw, str)
			break
		}
	}

	//msgREGEX := `` +
	/* Must start with :
	 * Can have N charactes, digits and underscore
	 * Must have one !
	 * Username
	**/
	//`(:)[\w+]{1,}(!)` +

	/* Can have N charactes, digits and underscore
	 * Followed by '@'
	 * Username
	**/
	//`[\w+]{1,}(@)` +

	/* Can have N charactes, digits and underscore
	 * Followed by '.tmi.twitch.tv'
	 * Username
	**/
	//`[\w+]{1,}(.tmi.twitch.tv)` +

	/* Can have N charactes, digits and underscore, followed by '#'
	 * Command
	**/
	//`[\w\s]{1,}(#)` +

	/* Can have N charactes, digits and underscore, followed by ':'
	 * Username
	**/
	//`[\w\s]{1,}(:)` +
	//// User message
	//`[\s\S]{1,}`
	//r := regexp.MustCompile(`(:)[\w+]{1,}(!)[\w+]{1,}(@)[\w+]{1,}(.tmi.twitch.tv )[\w]{1,}( #)[\w]{1,}( :)[\s\S]{1,}`)
	//r := regexp.MustCompile(msgREGEX)

	return nil
}
