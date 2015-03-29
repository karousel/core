package collection

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

func (store Store) Add(collection *Collection) error {
	err := store.Database.Insert(collection)

	return err
}

func (store Store) GetById(id int64) (Collection, error) {
	var collection Collection

	err := store.Database.SelectOne(&collection, "select * from collections where id=$1", id)

	return collection, err
}

func (store Store) GetAll() ([]Collection, error) {
	var collections []Collection

	_, err := store.Database.Select(&collections, "select * from collections order by id")

	return collections, err
}

func (store Store) Delete(collection *Collection) (int64, error) {
	count, err := store.Database.Delete(collection)

	return count, err
}

func (store Store) Update(collection *Collection) (int64, error) {
	count, err := store.Database.Update(collection)

	return count, err
}

func (store Store) Size() (int64, error) {
	size, err := store.Database.SelectInt("select count(*) from collections")

	return size, err
}
