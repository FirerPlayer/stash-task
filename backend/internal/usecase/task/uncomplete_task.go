package task

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type UncompleteTaskUseCase struct {
	TasksGateway gateway.TasksGateway
}

func NewUncompleteTaskUseCase(taskGateway gateway.TasksGateway) *UncompleteTaskUseCase {
	return &UncompleteTaskUseCase{
		TasksGateway: taskGateway,
	}
}

func (u *UncompleteTaskUseCase) Execute(ctx context.Context, input dto.UncompleteTaskInputDTO) error {
	task, err := u.TasksGateway.GetTaskByID(ctx, input.ID)
	if err != nil {
		return errors.New("failed to get task --> " + err.Error())
	}
	err = u.TasksGateway.UncompleteTaskByID(ctx, task.ID.String())
	if err != nil {
		return errors.New("failed to uncomplete task --> " + err.Error())
	}
	return nil
}
