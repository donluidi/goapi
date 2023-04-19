package database

import (
	"github.com/donluidi/goapi/internal/entity"
)

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}
