package database

import (
	"database/sql"

	"github.com/karousel/core/album"
	"github.com/karousel/core/collection"
	"github.com/karousel/core/file"
	"github.com/karousel/core/metadata"
	"github.com/karousel/core/photo"
	"github.com/karousel/core/token"
	"github.com/karousel/core/user"

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

	filesTable := databaseMap.AddTableWithName(file.File{}, "files")
	filesTable.SetKeys(true, "id")
	filesTable.ColMap("parent_id").SetNotNull(true)
	filesTable.ColMap("checksum").SetNotNull(true)

	metadataTable := databaseMap.AddTableWithName(metadata.Metadata{}, "metadata")
	metadataTable.SetKeys(true, "id")
	metadataTable.ColMap("parent_id").SetNotNull(true)
	metadataTable.ColMap("raw").SetNotNull(true)

	usersTable := databaseMap.AddTableWithName(user.User{}, "users")
	usersTable.SetKeys(true, "id")
	usersTable.ColMap("email_address").SetNotNull(true)
	usersTable.ColMap("email_address").SetUnique(true)
	usersTable.ColMap("password").SetNotNull(true)
	usersTable.ColMap("name").SetNotNull(true)
	usersTable.ColMap("is_administrator").SetNotNull(true)

	tokensTable := databaseMap.AddTableWithName(token.Token{}, "tokens")
	tokensTable.SetKeys(true, "id")
	tokensTable.ColMap("user_id").SetNotNull(true)
	tokensTable.ColMap("expires").SetNotNull(true)
	tokensTable.ColMap("token").SetNotNull(true)
	tokensTable.ColMap("token").SetUnique(true)

	err = databaseMap.CreateTablesIfNotExists()

	if err != nil {
		return databaseMap, err
	}

	return databaseMap, nil
}
