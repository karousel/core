package database

import (
	"github.com/karousel/core/album"
	"github.com/karousel/core/collection"
)

func (database Database) AddCollection(collection *collection.Collection) error {
	err := database.Connection.Insert(collection)

	return err
}

func (database Database) GetCollectionWithId(id int64) (collection.Collection, error) {
	var collection collection.Collection

	err := database.Connection.SelectOne(&collection, "select * from collections where id=$1", id)

	return collection, err
}

func (database Database) GetAllCollections() ([]collection.Collection, error) {
	var collections []collection.Collection

	_, err := database.Connection.Select(&collections, "select * from collections order by id")

	return collections, err
}

func (database Database) DeleteCollection(collection *collection.Collection) (int64, error) {
	count, err := database.Connection.Delete(collection)

	return count, err
}

func (database Database) UpdateCollection(collection *collection.Collection) (int64, error) {
	count, err := database.Connection.Update(collection)

	return count, err
}

func (database Database) CollectionCount() (int64, error) {
	count, err := database.Connection.SelectInt("select count(*) from collections")

	return count, err
}

func (database Database) GetAlbumsForCollection(collection *collection.Collection) ([]album.Album, error) {
	var albums []album.Album

	_, err := database.Connection.Select(&albums, "select * from albums where collection_id=$1", collection.Id)

	return albums, err
}
