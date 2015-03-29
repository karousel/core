package database

import (
	"database/sql"

	"github.com/karousel/core/collection"

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

	err = databaseMap.CreateTablesIfNotExists()

	if err != nil {
		return databaseMap, err
	}

	return databaseMap, nil
}
