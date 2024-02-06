package dto

import (
	"time"
)

// Create UserID
type CreateUserInputDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Password string `json:"password"`
}

type CreateUserOutputDTO struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Avatar    []byte    `json:"avatar"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// List All Users
type ListAllUsersInputDTO struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListAllUsersOutputDTO struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Avatar    []byte    `json:"avatar"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Update UserID
type UpdateUserInputDTO struct {
	UserID   string `json:"userID"`
	Email    string `json:"email"`
	Avatar   []byte `json:"avatar"`
	Bio      string `json:"bio"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Find UserID by ID
type GetUserByIDInputDTO struct {
	ID string `json:"id"`
}

type GetUserByIDOutputDTO struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Avatar    []byte    `json:"avatar"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Find UserID by Email
type GetUserByEmailInputDTO struct {
	Email string `json:"email"`
}
type GetUserByEmailOutputDTO struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Avatar    []byte    `json:"avatar"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Delete UserID by ID
type DeleteUserByIDInputDTO struct {
	ID string `json:"id"`
}

// Login
type LoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutputDTO struct {
	Token string `json:"token"`
}
