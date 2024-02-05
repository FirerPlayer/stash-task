package entity

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          uuid.UUID
	UserID      string
	Title       string
	Description string
	Priority    int
	CompletedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(user string, title, description string, priority int) *Task {
	if priority <= 0 {
		priority = 3
	}
	return &Task{
		ID:          uuid.New(),
		UserID:      user,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

}
