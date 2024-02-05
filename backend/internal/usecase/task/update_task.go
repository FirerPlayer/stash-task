package task

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/entity"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
	"time"
)

type UpdateTaskUseCase struct {
	TasksGateway gateway.TasksGateway
}

func NewUpdateTaskUseCase(tasksGateway gateway.TasksGateway) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		TasksGateway: tasksGateway,
	}
}

func (u *UpdateTaskUseCase) Execute(ctx context.Context, input dto.UpdateTaskInputDTO) error {
	err := u.TasksGateway.UpdateTaskByID(ctx, input.ID, &entity.Task{
		Title:       input.Title,
		Description: input.Description,
		Priority:    input.Priority,
		CompletedAt: input.CompletedAt,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return errors.New("failed to update task --> " + err.Error())
	}
	return nil

}
