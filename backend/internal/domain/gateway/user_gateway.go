package gateway

import (
	"context"
	"github.com/firerplayer/stash-task/backend/internal/domain/entity"
)

type UsersGateway interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	ListAllUsers(ctx context.Context, limit int, offset int) ([]*entity.User, error)
	DeleteUserByID(ctx context.Context, id string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUserByID(ctx context.Context, id string, user *entity.User) error
}
