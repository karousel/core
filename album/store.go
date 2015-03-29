package album

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

func (store Store) Add(album *Album) error {
	err := store.Database.Insert(album)

	return err
}

func (store Store) GetById(id int64) (Album, error) {
	var album Album

	err := store.Database.SelectOne(&album, "select * from albums where id=$1", id)

	return album, err
}

func (store Store) GetAll() ([]Album, error) {
	var albums []Album

	_, err := store.Database.Select(&albums, "select * from albums order by id")

	return albums, err
}

func (store Store) Delete(album *Album) (int64, error) {
	count, err := store.Database.Delete(album)

	return count, err
}

func (store Store) Update(album *Album) (int64, error) {
	count, err := store.Database.Update(album)

	return count, err
}

func (store Store) Size() (int64, error) {
	size, err := store.Database.SelectInt("select count(*) from albums")

	return size, err
}
