package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID                   uuid.UUID `json:"id"`
	Username             string    `json:"username"`
	Password             string    `json:"password"`
	PasswordConfirmation string    `json:"password_confirmation"`
}
