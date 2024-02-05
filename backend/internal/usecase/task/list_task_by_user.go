package task

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type ListTasksByUser struct {
	TasksGateway gateway.TasksGateway
}

func NewListTasksByUser(tasksGateway gateway.TasksGateway) *ListTasksByUser {
	return &ListTasksByUser{
		TasksGateway: tasksGateway,
	}
}

func (u *ListTasksByUser) Execute(ctx context.Context, input dto.ListAllTasksByUserIDInputDTO) ([]*dto.ListAllTasksByUserIDOutputDTO, error) {
	tasks, err := u.TasksGateway.ListTasksByUser(ctx, input.UserID, input.Limit, input.Offset)
	if err != nil {
		return nil, errors.New("failed to get tasks by user -> " + err.Error())
	}
	var output []*dto.ListAllTasksByUserIDOutputDTO
	for _, task := range tasks {
		output = append(output, &dto.ListAllTasksByUserIDOutputDTO{
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
