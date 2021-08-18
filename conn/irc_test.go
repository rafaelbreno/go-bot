package conn

import (
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/rafaelbreno/go-bot/test"
)

func TestIRC(t *testing.T) {
	tts := []test.TestCases{}

	{
		os.Setenv("IRC_URL", "irc://irc.chat.twitch.tv")
		os.Setenv("IRC_PORT", "6667")
		conn, err := NewIRC()
		connWant, _ := net.Dial("tcp", fmt.Sprintf(ircConnURL, os.Getenv("IRC_URL"), os.Getenv("IRC_PORT")))
		tts = append(tts, test.TestCases{
			Name:     "NewIRC OK - IRC",
			Got:      conn,
			Want:     connWant,
			TestType: test.Equal,
		})
		tts = append(tts, test.TestCases{
			Name:     "NewIRC OK - Error",
			Got:      err,
			TestType: test.Nil,
		})
	}

	{
		os.Setenv("IRC_URL", "some_random_url")
		os.Setenv("IRC_PORT", "not_even_a_port")
		conn, err := NewIRC()
		_, errWant := net.Dial("tcp", fmt.Sprintf(ircConnURL, os.Getenv("IRC_URL"), os.Getenv("IRC_PORT")))
		tts = append(tts, test.TestCases{
			Name:     "NewIRC OK - IRC",
			Got:      conn,
			Want:     &IRC{},
			TestType: test.Equal,
		})

		tts = append(tts, test.TestCases{
			Name:     "NewIRC Not - Error",
			Got:      err,
			Want:     errWant,
			TestType: test.Equal,
		})
	}

	test.RunTests(t, tts)
}
