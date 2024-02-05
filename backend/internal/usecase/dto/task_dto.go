package dto

import "time"

// Create Task
type CreateTaskInputDTO struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	UserID      string `json:"userID,omitempty"`
}

type CreateTaskOutputDTO struct {
	ID          string    `json:"id"`
	UserID      string    `json:"UserID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	CompletedAt time.Time `json:"completedAt,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

// Update Task
type UpdateTaskInputDTO struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    int       `json:"priority,omitempty"`
	CompletedAt time.Time `json:"completedAt,omitempty"`
}

// Get Task
type GetTaskByIDInputDTO struct {
	ID string `json:"id"`
}

type GetTaskByIDOutputDTO struct {
	ID          string    `json:"id"`
	UserID      string    `json:"UserID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	CompletedAt time.Time `json:"completedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// List All Tasks
type ListAllTasksInputDTO struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListAllTasksOutputDTO struct {
	ID          string    `json:"id"`
	UserID      string    `json:"UserID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	CompletedAt time.Time `json:"completedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// List All Tasks By UserID
type ListAllTasksByUserIDInputDTO struct {
	UserID string `json:"UserID,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Offset int    `json:"offset,omitempty"`
}

type ListAllTasksByUserIDOutputDTO struct {
	ID          string    `json:"id"`
	UserID      string    `json:"UserID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	CompletedAt time.Time `json:"completedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Delete Task
type DeleteTaskInputDTO struct {
	ID string `json:"id"`
}

// Complete Task
type CompleteTaskInputDTO struct {
	ID string `json:"id"`
}

// Uncomplete Task
type UncompleteTaskInputDTO struct {
	ID string `json:"id"`
}
