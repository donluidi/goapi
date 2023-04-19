package entity

import (
	"goapi/pkg"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       pkg.Id `json:"id"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"-"`
}

func NewUser(new_name, new_email, new_password string) *User {
	return &User{
		id:       pkg.NewId(),
		name:     new_name,
		email:    new_email,
		password: generateHashPassword(new_password),
	}
}

func generateHashPassword(new_password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

func (u *User) GetId() pkg.Id {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}
