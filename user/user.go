package user

import (
	"github.com/karousel/core/token"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id            int64  `db:"id"`
	Name          string `db:"name"`
	EmailAddress  string `db:"email_address"`
	Password      string `db:"password"`
	Administrator bool   `db:"is_administrator"`
}

func NewUser(name string, emailAddress string, password string, administrator bool) (User, error) {
	var user User

	user.Name = name
	user.EmailAddress = emailAddress
	user.Administrator = administrator

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return user, err
	}

	user.Password = string(hash)

	return user, nil
}

func (u User) AuthenticateWithPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return (err == nil)
}

func (u User) NewAuthenticationToken() (token.Token, error) {
	token, err := token.NewTokenForUserId(u.Id)

	if err != nil {
		return token, err
	}

	return token, nil
}
