package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CommandType define what kinda of parsing
// will be executed at one command.
type CommandType string

const (
	// Common commands, have a trigger
	// and an answer.
	Common CommandType = "COMMON"
	// Repeat commands, are those which
	// will be executed every X period of time.
	Repeat CommandType = "REPEAT"
)

// Command stores data related.
type Command struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Trigger     string    `json:"trigger" gorm:"size:16"`
	Template    string    `json:"template" gorm:"size:400"`
	Cooldown    string    `json:"cooldown" gorm:"size:16"`
	CommandType string    `json:"command_type" gorm:"type:varchar;size:16"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// ToJSON convert Command struct to
// a map type
func (c Command) ToJSON() CommandJSON {
	return CommandJSON{
		ID:          c.ID,
		Trigger:     c.Trigger,
		Template:    c.Template,
		CommandType: c.CommandType,
		Cooldown:    c.Cooldown,
	}
}

func (c *Command) UpdateFields(cmdFields Command) {
	if c.Trigger != cmdFields.Trigger {
		c.Trigger = cmdFields.Trigger
	}
	if c.Template != cmdFields.Template {
		c.Template = cmdFields.Template
	}
	if c.Cooldown != cmdFields.Cooldown {
		c.Cooldown = cmdFields.Cooldown
	}
	if c.CommandType != cmdFields.CommandType {
		c.CommandType = cmdFields.CommandType
	}
}

// CommandJSON DTO to receive data from
// http request
type CommandJSON struct {
	ID          uuid.UUID `json:"id"`
	Trigger     string    `json:"trigger"`
	Template    string    `json:"template"`
	Cooldown    string    `json:"cooldown"`
	CommandType string    `json:"command_type"`
}

func (c *CommandJSON) ToCommand() Command {
	return Command{
		ID:          c.ID,
		Trigger:     c.Trigger,
		Template:    c.Template,
		CommandType: c.CommandType,
		Cooldown:    c.Cooldown,
	}
}

// ToJSONString returns CommandJSON as JSON string
func (c *CommandJSON) ToJSONString() (string, error) {

	b, err := json.Marshal(c)

	return string(b), err
}
