package photo

import (
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

func (store Store) Add(photo *Photo) error {
	err := store.Database.Insert(photo)

	return err
}

func (store Store) GetById(id int64) (Photo, error) {
	var photo Photo

	err := store.Database.SelectOne(&photo, "select * from photos where id=$1", id)

	return photo, err
}

func (store Store) GetAll() ([]Photo, error) {
	var photos []Photo

	_, err := store.Database.Select(&photos, "select * from photos order by id")

	return photos, err
}

func (store Store) Delete(photo *Photo) (int64, error) {
	count, err := store.Database.Delete(photo)

	return count, err
}

func (store Store) Update(photo *Photo) (int64, error) {
	count, err := store.Database.Update(photo)

	return count, err
}

func (store Store) Size() (int64, error) {
	size, err := store.Database.SelectInt("select count(*) from photos")

	return size, err
}
