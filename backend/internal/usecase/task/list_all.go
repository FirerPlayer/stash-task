package task

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type ListAllTasksUseCase struct {
	TasksGateway gateway.TasksGateway
}

func NewListAllTasksUseCase(tasksGateway gateway.TasksGateway) *ListAllTasksUseCase {
	return &ListAllTasksUseCase{
		TasksGateway: tasksGateway,
	}
}

func (u *ListAllTasksUseCase) Execute(ctx context.Context, input dto.ListAllTasksInputDTO) ([]*dto.ListAllTasksOutputDTO, error) {
	tasks, err := u.TasksGateway.ListAllTasks(ctx, input.Limit, input.Offset)
	if err != nil {
		return nil, errors.New("failed to get tasks -> " + err.Error())
	}
	var output []*dto.ListAllTasksOutputDTO
	for _, task := range tasks {
		output = append(output, &dto.ListAllTasksOutputDTO{
			ID:          task.ID.String(),
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Priority:    task.Priority,
			CompletedAt: task.CompletedAt,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}
	return output, nil

}
