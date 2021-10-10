package user

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `pg:"type:uuid,default:gen_random_uuid()"`
	Username string
	Password string
}
