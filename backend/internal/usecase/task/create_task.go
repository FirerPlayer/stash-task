package task

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/entity"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type CreateTaskUseCase struct {
	TasksGateway gateway.TasksGateway
}

func NewCreateTaskUseCase(tasksGateway gateway.TasksGateway) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		TasksGateway: tasksGateway,
	}
}

func (u *CreateTaskUseCase) Execute(ctx context.Context, input dto.CreateTaskInputDTO) (*dto.CreateTaskOutputDTO, error) {
	newTask := entity.NewTask(input.UserID, input.Title, input.Description, input.Priority)
	tsk, err := u.TasksGateway.CreateTask(ctx, newTask)
	if err != nil {
		return nil, errors.New("failed to create task -> " + err.Error())
	}
	return &dto.CreateTaskOutputDTO{
		ID:          tsk.ID.String(),
		UserID:      tsk.UserID,
		Title:       tsk.Title,
		Description: tsk.Description,
		Priority:    tsk.Priority,
		CompletedAt: tsk.CompletedAt,
		CreatedAt:   tsk.CreatedAt,
		UpdatedAt:   tsk.UpdatedAt,
	}, nil

}
