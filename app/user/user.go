package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Username string `json:"username"`
	Coins    string `json:"coins"`
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

func GetByTag(username, tag string) (string, error) {
	user := users[username]

	if user == (User{}) {
		err := fmt.Errorf("Username %s not found!", username)
		log.Println(err)
		return "", err
	}

	r := reflect.ValueOf(user)
	f := reflect.Indirect(r).FieldByName(tag)

	log.Println(f)

	return string(f.String()), nil
}

func SaveUsers() {
	log.Println("Users saved!")
}
