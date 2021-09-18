package storage

import "github.com/rafaelbreno/go-bot/services/message-reader/command"

type Redis struct {
}

// NewRedis create a new Redis instance
func NewRedis() *Redis {
	return &Redis{}
}

// GetChannels retrieves from Redis a list
// of all channels
func (r *Redis) GetChannels(key string) []string {
	var channels []string

	return channels
}

// GetCommand retrieves from Redis a command
// from a given channel and key
func (r *Redis) GetCommand(channel, key string) command.Command {
	return command.Command{}
}
