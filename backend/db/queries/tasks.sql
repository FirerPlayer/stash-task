-- name: CreateTask :one
INSERT INTO tasks ("user",
                   title,
                   description,
                   priority,
                   completed_at,
                   created_at,
                   updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
returning *;
-- name: GetTaskByID :one
SELECT *
FROM tasks
WHERE id = $1;
-- name: ListAllTasks :many
SELECT *
FROM tasks
LIMIT $1 OFFSET $2;
-- name: UpdateTaskByID :exec
UPDATE tasks
SET title        = $1,
    description  = $2,
    priority     = $3,
    completed_at = $4,
    updated_at   = $5
WHERE id = $6;
-- name: DeleteTaskByID :exec
DELETE
FROM tasks
WHERE id = $1;
-- name: CompleteTaskByID :exec
UPDATE tasks
SET completed_at = $1
WHERE id = $2;
-- name: UncompleteTaskByID :exec
UPDATE tasks
SET completed_at = null
WHERE id = $1;
-- name: ListTasksByUser :many
SELECT *
FROM tasks
WHERE "user" = $1
LIMIT $2 OFFSET $3;