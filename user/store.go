package photo

import (
	"github.com/karousel/core/file"

	gorp "gopkg.in/gorp.v1"
)

type Store struct {
	Database *gorp.DbMap
}

func NewStore(database *gorp.DbMap) (Store, error) {
	store := Store{
		Database: database,
	}

	return store, nil
}

func (store Store) Add(user *User) error {
	err := store.Database.Insert(user)

	return err
}

func (store Store) GetById(id int64) (User, error) {
	var user User

	err := store.Database.SelectOne(&user, "select * from users where id=$1", id)

	return user, err
}

func (store Store) GetAll() ([]User, error) {
	var user []User

	_, err := store.Database.Select(&users, "select * from users order by id")

	return user, err
}

func (store Store) Delete(user *User) (int64, error) {
	count, err := store.Database.Delete(user)

	return count, err
}

func (store Store) Update(user *User) (int64, error) {
	count, err := store.Database.Update(user)

	return count, err
}

func (store Store) Size() (int64, error) {
	size, err := store.Database.SelectInt("select count(*) from users")

	return size, err
}
