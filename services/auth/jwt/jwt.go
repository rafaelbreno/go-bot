package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rafaelbreno/go-bot/auth/storage"
)

type JWT struct {
	Storage *storage.Storage
}

func NewJWT(sto *storage.Storage) *JWT {
	return &JWT{
		Storage: sto,
	}
}

var (
	MY_SIGN = []byte(`some_sign`)
)

type MyClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func (j *JWT) NewToken(id string) (string, error) {
	myClaim := MyClaim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

	return token.SignedString(MY_SIGN)
}

func (j *JWT) CheckToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return MY_SIGN, nil
	})

	if err != nil {
		return err
	}

	if token.Valid {
		return nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return fmt.Errorf("invalid token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return fmt.Errorf("expired token")
		} else {
			return err
		}
	}
	return err
}
