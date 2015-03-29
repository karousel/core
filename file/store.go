package file

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

func (store Store) Add(file *File) error {
	err := store.Database.Insert(file)

	return err
}

func (store Store) GetById(id int64) (File, error) {
	var file File

	err := store.Database.SelectOne(&file, "select * from files where id=$1", id)

	return file, err
}

func (store Store) GetAll() ([]File, error) {
	var files []File

	_, err := store.Database.Select(&files, "select * from files order by id")

	return files, err
}

func (store Store) Delete(file *File) (int64, error) {
	count, err := store.Database.Delete(file)

	return count, err
}

func (store Store) Update(file *File) (int64, error) {
	count, err := store.Database.Update(file)

	return count, err
}

func (store Store) Size() (int64, error) {
	size, err := store.Database.SelectInt("select count(*) from files")

	return size, err
}
