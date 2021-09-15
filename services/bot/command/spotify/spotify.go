package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/rafaelbreno/go-bot/internal"
)

type Player struct {
	Ctx *internal.Context
}

type SpotifyResp struct {
	Item Item `json:"item"`
}

type MusicResp struct {
	Items []Item `json:"items"`
}

type Item struct {
	Artists []Artist `json:"artists"`
	Albums  Album    `json:"album"`
	Name    string   `json:"name"`
	URI     string   `json:"uri"`
}

type Artist struct {
	Name string `json:"name"`
}

type Album struct {
	Name string `json:"name"`
}

func (p *Player) CurrentlyPlaying(token string) string {
	client := &http.Client{}

	// https://api.spotify.com/v1/me/player/currently-playing?market=BR
	// Headers
	// Accept: application/json
	// Content-Type: application/json
	// Authorization: Bearer <token>
	req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/me/player/currently-playing?market=BR", nil)

	if err != nil {
		p.Ctx.Logger.Error(err.Error())
		return "An error occurred"
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		p.Ctx.Logger.Error(err.Error())
		return "An error occurred"
	}

	resBody := res.Body
	resBytes, _ := ioutil.ReadAll(resBody)

	var s SpotifyResp

	if err := json.Unmarshal(resBytes, &s); err != nil {
		p.Ctx.Logger.Error(err.Error())
		return "An error occurred"
	}

	return fmt.Sprintf("%s - %s", s.Item.Name, s.Item.Artists[0].Name)
}

func (p *Player) AddMusic(music, token string) string {
	uri := p.getURI(music, token)

	if uri == nil {
		return "Music not found!"
	}

	client := &http.Client{}

	// https://api.spotify.com/v1/me/player/currently-playing?market=BR
	// Headers
	// Accept: application/json
	// Content-Type: application/json
	// Authorization: Bearer <token>
	req, err := http.NewRequest(http.MethodPost, "https://api.spotify.com/v1/me/player/queue", url.Values{
		"uri": {uri},
	}.Encode())

	if err != nil {
		p.Ctx.Logger.Error(err.Error())
		return "An error occurred"
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		p.Ctx.Logger.Error(err.Error())
		return "An error occurred"
	}

	resBody := res.Body
	resBytes, _ := ioutil.ReadAll(resBody)

	var s SpotifyResp

	if err := json.Unmarshal(resBytes, &s); err != nil {
		p.Ctx.Logger.Error(err.Error())
		return "An error occurred"
	}

	return fmt.Sprintf("%s - %s", uri.Name, uri.Albums.Name)
}

func (p *Player) getURI(music, token string) Item {
	client := &http.Client{}

	// https://api.spotify.com/v1/me/player/currently-playing?market=BR
	// Headers
	// Accept: application/json
	// Content-Type: application/json
	// Authorization: Bearer <token>
	req, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/search", url.Values{
		"q":      {"music"},
		"type":   {"track"},
		"market": {"BR"},
		"limit":  {"1"},
	}.Encode())

	if err != nil {
		p.Ctx.Logger.Error(err.Error())
		return Item{}
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		p.Ctx.Logger.Error(err.Error())
		return Item{}
	}

	resBody := res.Body
	resBytes, _ := ioutil.ReadAll(resBody)

	var s MusicResp

	if err := json.Unmarshal(resBytes, &s); err != nil {
		p.Ctx.Logger.Error(err.Error())
		return Item{}
	}

	if len(s.Items) == 0 {
		return Item{}
	}

	return s.Items[0]
}
