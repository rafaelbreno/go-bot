package utils

import (
	"regexp"
	"testing"

	"github.com/rafaelbreno/go-bot/test"
)

func TestRandom(t *testing.T) {
	tts := []test.TestCases{}

	{
		want := "A"
		got := Random([]string{"A", "B", "C", "D"}, "B", "C", "D")
		tts = append(tts, test.TestCases{
			TestType: test.Equal,
			Want:     want,
			Got:      got,
			Name:     "Helper Message Random",
		})
	}

	test.RunTests(t, tts)
}

func TestFind(t *testing.T) {
	tts := []test.TestCases{}

	{
		want := true
		got := Find([]string{"A", "B", "C", "D"}, "B")
		tts = append(tts, test.TestCases{
			TestType: test.Equal,
			Want:     want,
			Got:      got,
			Name:     "Helper Message Find",
		})
	}
	{
		want := false
		got := Find([]string{"A", "B", "C", "D"}, "X")
		tts = append(tts, test.TestCases{
			TestType: test.Equal,
			Want:     want,
			Got:      got,
			Name:     "Helper Message Find",
		})
	}

	test.RunTests(t, tts)
}

func TestReplace(t *testing.T) {
	tts := []test.TestCases{}

	repMap := map[string]string{
		"{user}": "foo",
	}

	{
		template := "some message {user}"
		want := "some message foo"
		got := Replace(template, repMap, map[string][]string{})
		tts = append(tts, test.TestCases{
			TestType: test.Equal,
			Want:     want,
			Got:      got,
			Name:     "Helper Message Replace",
		})
	}

	test.RunTests(t, tts)
}

func TestRandomRange(t *testing.T) {
	tts := []test.TestCases{}
	{
		got := RandomRange("some message {random.1-20}")
		wantReg := regexp.MustCompile(`some message [1,2]{1,}`)

		tts = append(tts, test.TestCases{
			TestType:  test.Regex,
			Got:       got,
			RegexRule: wantReg,
			Name:      "Helper Message RandomRange",
		})
	}

	test.RunTests(t, tts)
}
