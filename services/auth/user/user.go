package user

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `pg:"type:uuid,pk,default:gen_random_uuid()"`
	Username string    `pg:"username,unique"`
	Password string
}
