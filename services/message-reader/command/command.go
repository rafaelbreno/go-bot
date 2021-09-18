package command

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rafaelbreno/go-bot/services/message-reader/helpers"
)

// Command stores all data related
// to one command
type Command struct {
	Trigger string  `json:"trigger"`
	Answer  string  `json:"answer"`
	Fields  []Field `json:"fields"`
	SentBy  string  `json:"sent_by"`
}

type Field struct {
	Key       string   `json:"key"`
	Values    []string `json:"values"`
	Blacklist []string `json:"blacklist"`
}

func (c *Command) Parse() string {
	c.Replace("{sent_by}", c.SentBy)

	for _, field := range c.Fields {
		c.Replace(field.Key, c.GetRandom(field.Values, field.Blacklist))
	}

	c.CheckDefault()

	return c.Answer
}

func (c *Command) Replace(key, value string) {
	c.Answer = strings.ReplaceAll(c.Answer, key, value)
}

var (
	regexRandom = regexp.MustCompile(`\{random\.[0-9]{1,}\-[0-9]{1,}\}`)
	regexNumber = regexp.MustCompile("[0-9]{1,}")
)

// CheckDefault find and replace default fields
// e.g: {random.1-9}
func (c *Command) CheckDefault() {
	randomFields := regexRandom.FindAllString(c.Answer, -1)
	fmt.Println(randomFields)

	c.replaceRandomField(randomFields)
}

// replaceRandomField receives a list with fields
// to parse the max/min number and replace it
// in a string
func (c *Command) replaceRandomField(fields []string) {
	for _, k := range fields {
		nums := regexNumber.FindAllString(k, -1)
		minNum, _ := strconv.Atoi(nums[0])
		maxNum, _ := strconv.Atoi(nums[1])
		if minNum > maxNum {
			tmp := minNum
			minNum = maxNum
			maxNum = tmp
		}
		randNum := rand.Intn(maxNum) - minNum
		c.Answer = strings.ReplaceAll(c.Answer, k, strconv.Itoa(randNum))
	}
}

// GetRandom returns a value from a list
// checking if there's a value in a blacklist
func (c *Command) GetRandom(values, blacklist []string) (val string) {
	rand.Seed(time.Now().Unix())

	if len(blacklist) == 0 {
		val = values[rand.Intn(len(values))]
		return
	}

	for {
		val = values[rand.Intn(len(values))]
		if !helpers.Find(blacklist, val) {
			break
		}
	}

	return
}
