// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, avatar, username, password, bio, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
returning id, email, avatar, username, password, bio, created_at, updated_at
`

type CreateUserParams struct {
	Email     string
	Avatar    pgtype.Text
	Username  string
	Password  string
	Bio       pgtype.Text
	UpdatedAt pgtype.Timestamp
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.Avatar,
		arg.Username,
		arg.Password,
		arg.Bio,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Avatar,
		&i.Username,
		&i.Password,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE
FROM users
WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUserByID, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, avatar, username, password, bio, created_at, updated_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Avatar,
		&i.Username,
		&i.Password,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, avatar, username, password, bio, created_at, updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Avatar,
		&i.Username,
		&i.Password,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAllUsers = `-- name: ListAllUsers :many
SELECT id, email, avatar, username, password, bio, created_at, updated_at
FROM users
LIMIT $1 OFFSET $2
`

type ListAllUsersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListAllUsers(ctx context.Context, arg ListAllUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listAllUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Avatar,
			&i.Username,
			&i.Password,
			&i.Bio,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserByID = `-- name: UpdateUserByID :exec
UPDATE users
SET email      = $1,
    avatar     = $2,
    username   = $3,
    password   = $4,
    updated_at = $5
WHERE id = $6
`

type UpdateUserByIDParams struct {
	Email     string
	Avatar    pgtype.Text
	Username  string
	Password  string
	UpdatedAt pgtype.Timestamp
	ID        pgtype.UUID
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) error {
	_, err := q.db.Exec(ctx, updateUserByID,
		arg.Email,
		arg.Avatar,
		arg.Username,
		arg.Password,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}