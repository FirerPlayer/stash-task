package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Avatar    []byte
	Username  string
	Password  string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(email, bio, username, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Email:     email,
		Username:  username,
		Bio:       bio,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
