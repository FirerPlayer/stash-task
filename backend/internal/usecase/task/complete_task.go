package task

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type CompletTaskUseCase struct {
	TasksGateway gateway.TasksGateway
}

func NewCompletTaskUseCase(tasksGateway gateway.TasksGateway) *CompletTaskUseCase {
	return &CompletTaskUseCase{
		TasksGateway: tasksGateway,
	}
}

func (u *CompletTaskUseCase) Execute(ctx context.Context, input dto.CompleteTaskInputDTO) error {
	task, err := u.TasksGateway.GetTaskByID(ctx, input.ID)
	if err != nil {
		return errors.New("failed to get task --> " + err.Error())
	}
	err = u.TasksGateway.CompleteTaskByID(ctx, task.ID.String())
	if err != nil {
		return errors.New("failed to complete task --> " + err.Error())
	}
	return nil
}
