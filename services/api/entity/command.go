package entity

import (
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
	ID          uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid()"`
	Trigger     string      `gorm:"size:16"`
	Template    string      `gorm:"size:400"`
	Cooldown    string      `gorm:"size:16"`
	CommandType CommandType `gorm:"type:command_type"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
