package user

type User struct {
	Id            int64  `db:"id"`
	Name          string `db:"name"`
	EmailAddress  string `db:"email_address"`
	Password      string `db:"password"`
	Administrator bool   `db:"is_administrator"`
}
