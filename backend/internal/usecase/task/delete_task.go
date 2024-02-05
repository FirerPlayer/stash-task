package task

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type DeleteTaskUseCase struct {
	TasksGateway gateway.TasksGateway
}

func NewDeleteTaskUseCase(tasksGateway gateway.TasksGateway) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		TasksGateway: tasksGateway,
	}
}

func (u *DeleteTaskUseCase) Execute(ctx context.Context, input dto.DeleteTaskInputDTO) error {
	err := u.TasksGateway.DeleteTaskByID(ctx, input.ID)
	if err != nil {
		return errors.New("error deleting task --> " + err.Error())
	}
	return nil

}
