package task

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type GetTaskByIDUseCase struct {
	TasksGateway gateway.TasksGateway
}

func NewGetTaskByIDUseCase(tasksGateway gateway.TasksGateway) *GetTaskByIDUseCase {
	return &GetTaskByIDUseCase{
		TasksGateway: tasksGateway,
	}
}

func (u *GetTaskByIDUseCase) Execute(ctx context.Context, input dto.GetTaskByIDInputDTO) (*dto.GetTaskByIDOutputDTO, error) {
	task, err := u.TasksGateway.GetTaskByID(ctx, input.ID)
	if err != nil {
		return nil, errors.New("failed to get task -> " + err.Error())
	}
	return &dto.GetTaskByIDOutputDTO{
		ID:          task.ID.String(),
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CompletedAt: task.CompletedAt,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}
