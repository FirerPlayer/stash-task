package usecase

import (
	"context"
	"errors"
	"github.com/firerplayer/stash-task/backend/internal/domain/gateway"
	"github.com/firerplayer/stash-task/backend/internal/usecase/dto"
)

type FindUserByEmailUseCase struct {
	UsersGateway gateway.UsersGateway
}

func NewFindUserByEmailUseCase(usersGateway gateway.UsersGateway) *FindUserByEmailUseCase {
	return &FindUserByEmailUseCase{
		UsersGateway: usersGateway,
	}
}

// Execute retrieves a user by email.
//
// ctx is the context of the request.
// email is the email of the user to retrieve.
// Returns a pointer to a userID and an error if any occurred.
func (uc *FindUserByEmailUseCase) Execute(ctx context.Context, email string) (*dto.GetUserByEmailOutputDTO, error) {
	user, err := uc.UsersGateway.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("failed to find user -> " + err.Error())
	}

	return &dto.GetUserByEmailOutputDTO{
		ID:        user.ID.String(),
		Username:  user.Username,
		Avatar:    user.Avatar,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
