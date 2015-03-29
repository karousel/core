package metadata

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

func (store Store) Add(metadata *Metadata) error {
	err := store.Database.Insert(metadata)

	return err
}

func (store Store) GetById(id int64) (Metadata, error) {
	var metadata Metadata

	err := store.Database.SelectOne(&metadata, "select * from metadata where id=$1", id)

	return metadata, err
}

func (store Store) GetAll() ([]Metadata, error) {
	var metadata []Metadata

	_, err := store.Database.Select(&metadata, "select * from metadata order by id")

	return metadata, err
}

func (store Store) Delete(metadata *Metadata) (int64, error) {
	count, err := store.Database.Delete(metadata)

	return count, err
}

func (store Store) Update(metadata *Metadata) (int64, error) {
	count, err := store.Database.Update(metadata)

	return count, err
}

func (store Store) Size() (int64, error) {
	size, err := store.Database.SelectInt("select count(*) from metadata")

	return size, err
}
