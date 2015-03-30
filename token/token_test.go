package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserTokenCreation(t *testing.T) {
	assert := assert.New(t)
	var userId int64

	userId = 1

	token, err := NewTokenForUserId(userId)

	assert.Nil(err)

	assert.Equal(userId, token.UserId)
	assert.True(token.Expires)
	assert.Equal(64, len([]byte(token.Token)))
}
