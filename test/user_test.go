package test

import (
	"testing"

	"github.com/donluidi/goapi/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := entity.NewUser("John Doe", "j@j.com", "123456")
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.GetId())
	assert.NotEmpty(t, user.GetEmail())
	assert.Equal(t, "John Doe", user.GetName())
}

func TestUser_ValidatePassword(t *testing.T) {
	user := entity.NewUser("John Doe", "j@j.com", "123456")
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("123457"))
	assert.NotEqual(t, "123456", user.GetPassword())
}
