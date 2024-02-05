-- name: CreateUser :one
INSERT INTO users (email, avatar, username, password, bio, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
returning *;
-- name: ListAllUsers :many
SELECT *
FROM users
LIMIT $1 OFFSET $2;
-- name: UpdateUserByID :exec
UPDATE users
SET email      = $1,
    avatar     = $2,
    username   = $3,
    password   = $4,
    updated_at = $5
WHERE id = $6;
-- name: DeleteUserByID :exec
DELETE
FROM users
WHERE id = $1;
-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;
-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;