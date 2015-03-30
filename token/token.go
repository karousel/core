package token

import (
	"crypto/rand"
	"time"
)

type Token struct {
	Id         int64     `db:"id"`
	UserId     int64     `db:"user_id"`
	Expires    bool      `db:"expires"`
	Token      string    `db:"token"`
	Expiration time.Time `db:"expiration"`
}

func NewTokenForUserId(user int64) (Token, error) {
	var token Token
	size := 64

	value := make([]byte, size)
	_, err := rand.Read(value)

	if err != nil {
		return token, err
	}

	token.UserId = user
	token.Token = string(value)
	token.Expires = true
	token.Expiration = time.Now().UTC().Add(time.Hour * 72)

	return token, nil
}
