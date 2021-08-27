package internal

import (
	"testing"

	"github.com/rafaelbreno/go-bot/test"
	"go.uber.org/zap"
)

func TestContext(t *testing.T) {
	tts := []test.TestCases{}

	{
		l := zap.NewExample()
		chName := []string{"foo", "bar"}
		authToken := "bar"
		botName := "foobot"

		got := WriteContexts(l, authToken, botName, chName)

		want := map[string]*Context{}
		want["foo"] = &Context{
			Logger:      l,
			ChannelName: "foo",
			OAuthToken:  authToken,
			BotName:     botName,
		}
		want["bar"] = &Context{
			Logger:      l,
			ChannelName: "bar",
			OAuthToken:  authToken,
			BotName:     botName,
		}

		tts = append(tts, test.TestCases{
			Name:     "WriteContexts",
			Want:     want,
			Got:      got,
			TestType: test.Equal,
		})
	}

	test.RunTests(t, tts)
}
