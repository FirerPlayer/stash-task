package usecase

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type ListAllUsersUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewListAllUsersUseCase(usersGateway gateway.UsersGateway) *ListAllUsersUseCase {
	return &ListAllUsersUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *ListAllUsersUseCase) Execute(ctx context.Context, input dto.ListAllUsersInputDTO) ([]*dto.ListAllUsersOutputDTO, error) {
	users, err := uc.UsersGateway.ListAllUsers(ctx, input.Limit, input.Offset)
	if err != nil {
		return nil, errors.New("Failed to list users -> " + err.Error())
	}
	var output []*dto.ListAllUsersOutputDTO
	for _, user := range users {
		output = append(output, &dto.ListAllUsersOutputDTO{
			ID:        user.ID.String(),
			Username:  user.Username,
			Avatar:    user.Avatar,
			Bio:       user.Bio,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return output, nil
}
