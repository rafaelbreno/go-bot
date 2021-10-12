package command

import (
	"regexp"
	"testing"

	"github.com/rafaelbreno/go-bot/services/message-reader/test"
)

func TestCommand(t *testing.T) {
	tts := []test.TestCases{}

	{
		cmd := Command{
			Trigger: "!test",
			Answer:  "Hello {sent_by}",
			SentBy:  "foo",
		}

		tts = append(tts, test.TestCases{
			Name:     "Parse Hello",
			Want:     "Hello foo",
			Got:      cmd.Parse(),
			TestType: test.Equal,
		})
	}

	{
		cmd := Command{
			Trigger: "!test",
			Answer:  "Hello {sent_by}, you're {hero}",
			SentBy:  "foo",
			Fields: []Field{
				{
					Key:       "{hero}",
					Values:    []string{"A", "B", "C", "D", "E"},
					Blacklist: []string{},
				},
			},
		}

		tts = append(tts, test.TestCases{
			Name:      "Parse Hello",
			Got:       cmd.Parse(),
			RegexRule: regexp.MustCompile(`(Hello foo, you're) (A|B|C|D|E)`),
			TestType:  test.Regex,
		})
	}

	{
		cmd := Command{
			Trigger: "!test",
			Answer:  "Hello {sent_by}, you're {hero}",
			SentBy:  "foo",
			Fields: []Field{
				{
					Key:       "{hero}",
					Values:    []string{"A", "B", "C", "D", "E"},
					Blacklist: []string{"A", "C"},
				},
			},
		}

		tts = append(tts, test.TestCases{
			Name:      "Parse Hello",
			Got:       cmd.Parse(),
			RegexRule: regexp.MustCompile(`(Hello foo, you're) (B|D|E)`),
			TestType:  test.Regex,
		})
	}

	{
		cmd := Command{
			Trigger: "!test",
			Answer:  "Hello {sent_by}, you're {hero} {random.1-10}",
			SentBy:  "foo",
			Fields: []Field{
				{
					Key:       "{hero}",
					Values:    []string{"A", "B", "C", "D", "E"},
					Blacklist: []string{"A", "C"},
				},
			},
		}

		tts = append(tts, test.TestCases{
			Name:      "Parse Hello",
			Got:       cmd.Parse(),
			RegexRule: regexp.MustCompile(`(Hello foo, you're) (B|D|E) [1-9]{1,2}`),
			TestType:  test.Regex,
		})
	}

	test.RunTests(t, tts)
}
