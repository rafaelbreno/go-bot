package storage

import (
	"fmt"

	command "github.com/rafaelbreno/go-bot/services/message-reader/command"
)

// Storage consumes a in-memory database
// redis, Memcached, DynamoDB
type Storage interface {
	GetChannels(key string) []string
	GetCommand(key string) command.Command
}

const (
	channelsKey = "[channels]"
	commandKey  = "[%s][%s]"
)

// GetChannels retrieve channels list
// from a Storage
func GetChannels(s Storage) []string {
	return s.GetChannels(channelsKey)
}

// GetCommand retrieve command from Redis
func GetCommand(channel, key string, s Storage) command.Command {
	if channel == "" || key == "" {
		return command.Command{}
	}

	return s.GetCommand(fmt.Sprintf(commandKey, channel, key))
}
