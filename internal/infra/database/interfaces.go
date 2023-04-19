package database

import (
	"goapi/internal/entity"
)

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
