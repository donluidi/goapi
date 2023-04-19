package test

import (
	"testing"

	"github.com/donluidi/goapi/internal/entity"
	"github.com/donluidi/goapi/internal/infra/database"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user := entity.NewUser("John Doe", "j@j.com", "123456")
	userDB := database.NewUserDB(db)

	err = userDB.CreateUser(user)

	assert.Nil(t, err)
}
