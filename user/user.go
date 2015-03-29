package user

import (
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
