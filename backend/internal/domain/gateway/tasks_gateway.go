package gateway

import (
	"context"
	"github.com/firerplayer/stash-task/backend/internal/domain/entity"
)

type TasksGateway interface {
	CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error)
	GetTaskByID(ctx context.Context, id string) (*entity.Task, error)
	ListAllTasks(ctx context.Context, limit, offset int) ([]*entity.Task, error)
	ListTasksByUser(ctx context.Context, userID string, limit, offset int) ([]*entity.Task, error)
	CompleteTaskByID(ctx context.Context, id string) error
	UncompleteTaskByID(ctx context.Context, id string) error
	DeleteTaskByID(ctx context.Context, id string) error
	UpdateTaskByID(ctx context.Context, id string, task *entity.Task) error
}
