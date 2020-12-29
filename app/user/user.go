package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Username string `json:"username"`
	Coins    int    `json:"coins"`
	LastSeen string `json:"last_seen"`
}

var users map[string]User

func ImportUsers() {
	users = make(map[string]User)

	tempUsers := Users{}

	jsonFile, err := os.Open("db/users.json")

	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(jsonBytes, &tempUsers); err != nil {
		log.Fatal(err)
	}

	for _, val := range tempUsers.Users {
		users[val.Username] = val
	}
}

func GetUser(username string) {
	log.Println(users[username])
}

func SaveUsers() {
	log.Println("Users saved!")
}
