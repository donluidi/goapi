package test

import (
	"goapi/internal/entity"
	"goapi/internal/infra/database"
	"testing"

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
