package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	datasource = "dbname=karousel_test sslmode=disable"
)

func TestNewDatabase(t *testing.T) {
	assert := assert.New(t)

	database, err := NewDatabase(datasource)

	assert.Nil(err)

	err = database.Connection.Db.Ping()

	assert.Nil(err)
}
