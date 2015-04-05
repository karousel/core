package database

import (
	"testing"

	"github.com/karousel/core/collection"

	"github.com/stretchr/testify/assert"
)

var (
	database, _ = NewDatabase(datasource)
)

func TestAddCollection(t *testing.T) {
	assert := assert.New(t)

	count, err := database.CollectionCount()

	assert.Nil(err)

	var c collection.Collection
	c.Name = "My Collection"

	err = database.AddCollection(&c)

	assert.Nil(err)

	assert.NotEqual(c.Id, 0)

	newCount, err := database.CollectionCount()

	assert.Nil(err)

	assert.Equal(count+1, newCount)
}

func TestGetCollectionWithId(t *testing.T) {
	assert := assert.New(t)

	c, err := database.GetCollectionWithId(1)

	assert.Nil(err)

	assert.Equal(c.Id, 1)
	assert.Equal(c.Name, "My Collection")
}

func TestGetAllCollections(t *testing.T) {
	assert := assert.New(t)

	collections, err := database.GetAllCollections()

	assert.Nil(err)

	count, err := database.CollectionCount()

	assert.Nil(err)

	assert.Equal(len(collections), count)
}

func TestUpdateCollection(t *testing.T) {
	assert := assert.New(t)

	c, err := database.GetCollectionWithId(1)

	c.Name = "Your Collection"

	count, err := database.UpdateCollection(&c)

	assert.Equal(count, 1)
	assert.Nil(err)

	c, err = database.GetCollectionWithId(1)

	assert.Equal(c.Name, "Your Collection")
}

func TestDeleteCollection(t *testing.T) {
	assert := assert.New(t)

	c, err := database.GetCollectionWithId(1)

	assert.Nil(err)

	count, err := database.DeleteCollection(&c)

	assert.Nil(err)
	assert.Equal(count, 1)

	count, err = database.CollectionCount()

	assert.Nil(err)
	assert.Equal(count, 0)
}
