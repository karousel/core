package database

import (
	"database/sql"

	"github.com/karousel/core/album"
	"github.com/karousel/core/collection"
	"github.com/karousel/core/photo"

	_ "github.com/lib/pq"
	gorp "gopkg.in/gorp.v1"
)

func NewDatabase(datasource string) (*gorp.DbMap, error) {
	var databaseMap *gorp.DbMap

	database, err := sql.Open("postgres", datasource)

	if err != nil {
		return databaseMap, err
	}

	databaseMap = &gorp.DbMap{Db: database, Dialect: gorp.PostgresDialect{}}

	collectionsTable := databaseMap.AddTableWithName(collection.Collection{}, "collections")
	collectionsTable.SetKeys(true, "id")
	collectionsTable.ColMap("name").SetNotNull(true)
	collectionsTable.ColMap("name").SetUnique(true)

	albumsTable := databaseMap.AddTableWithName(album.Album{}, "albums")
	albumsTable.SetKeys(true, "id")
	albumsTable.ColMap("collection_id").SetNotNull(true)
	albumsTable.ColMap("name").SetNotNull(true)
	albumsTable.ColMap("name").SetUnique(true)

	photosTable := databaseMap.AddTableWithName(photo.Photo{}, "photos")
	photosTable.SetKeys(true, "id")
	photosTable.ColMap("album_id").SetNotNull(true)
	photosTable.ColMap("name").SetNotNull(true)
	photosTable.ColMap("uploaded").SetNotNull(true)
	photosTable.ColMap("status").SetNotNull(true)

	err = databaseMap.CreateTablesIfNotExists()

	if err != nil {
		return databaseMap, err
	}

	return databaseMap, nil
}
