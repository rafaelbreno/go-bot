package helpers

import (
	"testing"

	"github.com/rafaelbreno/go-bot/services/message-reader/test"
)

func TestSlice(t *testing.T) {
	tts := []test.TestCases{}

	{
		sl := []string{"A", "B", "C", "D", "E"}
		want := true
		got := FindInSliceStr(sl, "A")
		tts = append(tts, test.TestCases{
			Name:     "FindInSliceStr",
			Want:     want,
			Got:      got,
			TestType: test.Equal,
		})
	}
	{
		sl := []string{"A", "B", "C", "D", "E"}
		want := false
		got := FindInSliceStr(sl, "X")
		tts = append(tts, test.TestCases{
			Name:     "FindInSliceStr",
			Want:     want,
			Got:      got,
			TestType: test.Equal,
		})
	}
	{
		sl := []string{"A", "B", "C", "D", "E"}
		want := []string{"A", "B", "C", "D", "E"}
		got := RemoveElementStr(sl, "X")
		tts = append(tts, test.TestCases{
			Name:     "FindInSliceStr",
			Want:     want,
			Got:      got,
			TestType: test.Equal,
		})
	}
	{
		sl := []string{"A", "B", "C", "D", "E"}
		want := []string{"A", "B", "D", "E"}
		got := RemoveElementStr(sl, "C")
		tts = append(tts, test.TestCases{
			Name:     "FindInSliceStr",
			Want:     want,
			Got:      got,
			TestType: test.Equal,
		})
	}

	test.RunTests(t, tts)
}
