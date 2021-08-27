package conn

import (
	"os"
	"testing"

	"github.com/rafaelbreno/go-bot/internal"
	"github.com/rafaelbreno/go-bot/test"
)

func TestIRC(t *testing.T) {
	tts := []test.TestCases{}

	{
		os.Setenv("IRC_URL", "irc://irc.chat.twitch.tv")
		os.Setenv("IRC_PORT", "6667")
		conn, err := NewIRC(&internal.Context{})
		//netConn, _ := net.Dial("tcp", fmt.Sprintf(ircConnURL, os.Getenv("IRC_URL"), os.Getenv("IRC_PORT")))
		connWant := &IRC{
			Ctx: &internal.Context{},
			Msg: make(chan string, 1),
		}
		tts = append(tts, test.TestCases{
			Name:     "NewIRC OK - IRC",
			Got:      conn.Ctx,
			Want:     connWant.Ctx,
			TestType: test.Equal,
		})
		tts = append(tts, test.TestCases{
			Name:     "NewIRC OK - Error",
			Got:      err,
			TestType: test.Nil,
		})
	}

	test.RunTests(t, tts)
}
