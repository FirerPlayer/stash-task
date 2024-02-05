package usecase

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type GetUserByIDUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewGetUserByIdUseCase(usersGateway gateway.UsersGateway) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{
		UsersGateway: usersGateway,
	}
}

func (uc *GetUserByIDUseCase) Execute(ctx context.Context, id string) (*dto.GetUserByIDOutputDTO, error) {
	user, err := uc.UsersGateway.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.New("failed to find user -> " + err.Error())
	}

	return &dto.GetUserByIDOutputDTO{
		ID:        user.ID.String(),
		Username:  user.Username,
		Avatar:    user.Avatar,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
