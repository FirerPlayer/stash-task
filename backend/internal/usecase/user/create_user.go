package usecase

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/entity"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type CreateUserUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewCreateUserUseCase(userGateway gateway.UsersGateway) *CreateUserUseCase {
	return &CreateUserUseCase{
		UsersGateway: userGateway,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, input dto.CreateUserInputDTO) (*dto.CreateUserOutputDTO, error) {

	newUser, err := entity.NewUser(input.Email, input.Bio, input.Username, input.Password)
	if err != nil {
		return nil, errors.New("failed to create user --> " + err.Error())
	}
	createdUser, err := u.UsersGateway.CreateUser(ctx, newUser)
	if err != nil {
		return nil, errors.New("failed to create user --> " + err.Error())
	}

	return &dto.CreateUserOutputDTO{
		ID:        createdUser.ID.String(),
		Username:  createdUser.Username,
		Avatar:    createdUser.Avatar,
		Bio:       createdUser.Bio,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}
