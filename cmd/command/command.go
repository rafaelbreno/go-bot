package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Command struct {
	Identifier string
	Message    string
	Timeout    int
	Variables  []Variable
}

type Variable struct {
	Name  string
	Type  string
	Value interface{}
}

var commands []Command

func GetCommands() {
	parseCommands()
	fmt.Println(commands)
}

func parseCommands() {
	bytes := getCommandsByte()
	err := json.Unmarshal(bytes, &commands)

	if err != nil {
		panic(err)
	}
}

func getCommandsByte() []byte {
	json, err := ioutil.ReadFile("bd/commands.json")
	if err != nil {
		panic(err)
	}
	return json
}
