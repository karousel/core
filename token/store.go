package token

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

func (store Store) Add(token *Token) error {
	err := store.Database.Insert(token)

	return err
}

func (store Store) GetById(id int64) (Token, error) {
	var token Token

	err := store.Database.SelectOne(&token, "select * from tokens where id=$1", id)

	return token, err
}

func (store Store) GetByToken(value string) (Token, error) {
	var token Token

	err := store.Database.SelectOne(&token, "select * from tokens where token=$1", value)

	return token, err
}

func (store Store) GetAll() ([]Token, error) {
	var tokens []Token

	_, err := store.Database.Select(&tokens, "select * from tokens order by id")

	return tokens, err
}

func (store Store) Delete(token *Token) (int64, error) {
	count, err := store.Database.Delete(token)

	return count, err
}

func (store Store) Update(token *Token) (int64, error) {
	count, err := store.Database.Update(token)

	return count, err
}

func (store Store) Size() (int64, error) {
	size, err := store.Database.SelectInt("select count(*) from tokens")

	return size, err
}
