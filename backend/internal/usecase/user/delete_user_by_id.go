package usecase

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type DeleteUserByIDUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewDeleteUserByIDUseCase(usersGateway gateway.UsersGateway) *DeleteUserByIDUseCase {
	return &DeleteUserByIDUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *DeleteUserByIDUseCase) Execute(ctx context.Context, input dto.DeleteUserByIDInputDTO) error {
	err := uc.UsersGateway.DeleteUserByID(ctx, input.ID)
	if err != nil {
		return errors.New("Failed to delete user -> " + err.Error())
	}
	return nil
}
