package command

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/rafaelbreno/go-bot/internal"
)

type TMIViewers struct {
	Chatters struct {
		Broadcaster []string `json:"broadcaster"`
		Vips        []string `json:"vips"`
		Moderators  []string `json:"moderators"`
		Viewers     []string `json:"viewers"`
	} `json:"chatters"`
}

type CmdHelper struct {
	ctx     *internal.Context
	userURL string
}

func NewCMDHelper(ctx *internal.Context) *CmdHelper {
	return &CmdHelper{
		ctx:     ctx,
		userURL: fmt.Sprintf("https://tmi.twitch.tv/group/user/%s/chatters", ctx.ChannelName),
	}
}

func (c *CmdHelper) fetchUserList() []string {
	v := TMIViewers{}
	resp, _ := http.Get(c.userURL)
	_ = json.NewDecoder(resp.Body).Decode(&v)

	var list []string
	list = append(list, v.Chatters.Broadcaster...)
	list = append(list, v.Chatters.Viewers...)
	list = append(list, v.Chatters.Vips...)
	list = append(list, v.Chatters.Moderators...)
	return list
}

func random(list []string, blackList ...string) string {
	rand.Seed(time.Now().Unix())

	if len(blackList) == 0 {
		return list[rand.Intn(len(list))]
	}

	item := ""
	for {
		item = list[rand.Intn(len(list))]
		if !find(blackList, item) {
			break
		}
	}
	return item
}

func replace(str string, repMap keyMap) string {
	for key, val := range repMap {
		str = strings.ReplaceAll(str, key, val)
	}
	return str
}

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
