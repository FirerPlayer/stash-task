package usecase

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/entity"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
	"time"
)

type UpdateUserByIDUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewUpdateUserUseCase(usersGateway gateway.UsersGateway) *UpdateUserByIDUseCase {
	return &UpdateUserByIDUseCase{
		UsersGateway: usersGateway,
	}
}

// Execute updates the user with the given ID according to the provided user data.
//
// ctx is the context that the function executes under.
// input is the input data for updating the user.
// Returns an error if the update operation fails.
func (uc *UpdateUserByIDUseCase) Execute(ctx context.Context, input dto.UpdateUserInputDTO) error {
	err := uc.UsersGateway.UpdateUserByID(ctx, input.UserID, &entity.User{
		Email:     input.Email,
		Avatar:    input.Avatar,
		Username:  input.Username,
		Password:  input.Password,
		Bio:       input.Bio,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return errors.New("Failed to update user -> " + err.Error())
	}
	return nil
}
