package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

type Credentials struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
}

var oauth2Config *clientcredentials.Config

const base_url = "https://id.twitch.tv/oauth2/authorize?"

func getCredentials() Credentials {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cred := Credentials{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	return cred
}

func getAccessToken(cred Credentials) string {
	oauth2Config = &clientcredentials.Config{
		ClientID:     cred.ClientID,
		ClientSecret: cred.ClientSecret,
		TokenURL:     cred.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	return token.AccessToken
}

func clientCred() {
	cred := getCredentials()
	token := getAccessToken(cred)
	fmt.Println(token)
}

func main() {
	clientCred()
}
