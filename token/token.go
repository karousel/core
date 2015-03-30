package token

import (
	"time"
)

type Token struct {
	Id         int64     `db:"id"`
	UserId     int64     `db:"user_id"`
	Expires    bool      `db:"expires"`
	Token      string    `db:"token"`
	Expiration time.Time `db:"expiration"`
}
