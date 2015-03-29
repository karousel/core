package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	assert := assert.New(t)

	name := "Archibald Haddock"
	emailAddress := "haddock@marlinspike.com"
	password := "lochlomond"
	administrator := false

	user, err := NewUser(name, emailAddress, password, administrator)

	assert.Nil(err)
	assert.Equal(name, user.Name)
	assert.Equal(emailAddress, user.EmailAddress)
	assert.Equal(administrator, user.Administrator)
}
