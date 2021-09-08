package test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestType int

const (
	Equal TestType = iota
	Nil
	NotNil
	Regex
)

type TestCases struct {
	Name      string
	Want      interface{}
	Got       interface{}
	RegexRule *regexp.Regexp
	TestType  TestType
}

func RunTests(t *testing.T, tts []TestCases) {
	for _, tt := range tts {
		t.Run(tt.Name, func(t *testing.T) {
			switch tt.TestType {
			case Equal:
				assert.Equal(t, tt.Want, tt.Got)
			case Regex:
				assert.Regexp(t, tt.RegexRule, tt.Got)
			case Nil:
				assert.Nil(t, tt.Got)
			case NotNil:
				assert.NotNil(t, tt.Got)
			default:
			}
		})
	}
}
