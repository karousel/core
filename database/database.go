package database

import (
	"database/sql"

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

	return databaseMap, nil
}
