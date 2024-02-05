package repository

import (
	"context"
	"github.com/firerplayer/stash-task/backend/internal/domain/entity"
	pg "github.com/firerplayer/stash-task/backend/internal/infra/pg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepositoryPg struct {
	DB      *pgx.Conn
	Queries *pg.Queries
}

func NewUserRepositoryPg(dbt *pgx.Conn) *UserRepositoryPg {
	return &UserRepositoryPg{
		DB:      dbt,
		Queries: pg.New(dbt),
	}
}

func (r *UserRepositoryPg) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	params := pg.CreateUserParams{
		Email:     user.Email,
		Avatar:    pgtype.Text{string(user.Avatar), true},
		Username:  user.Username,
		Password:  user.Password,
		Bio:       pgtype.Text{String: user.Bio, Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: user.UpdatedAt, Valid: true},
	}

	createdUser, err := r.Queries.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	var createdUserEntity entity.User
	err = HydrateUser(createdUser, &createdUserEntity)
	if err != nil {
		return nil, err
	}
	return &createdUserEntity, nil
}

func (r *UserRepositoryPg) ListAllUsers(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	if limit == 0 {
		limit = 20
	}
	usersDb, err := r.Queries.ListAllUsers(ctx, pg.ListAllUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	var users []*entity.User
	for _, userDb := range usersDb {
		var user entity.User
		err := HydrateUser(userDb, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *UserRepositoryPg) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	userDb, err := r.Queries.GetUserByID(ctx, StringToUUID(id))
	if err != nil {
		return nil, err
	}
	var user entity.User
	err = HydrateUser(userDb, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryPg) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	userDb, err := r.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	var user entity.User
	err = HydrateUser(userDb, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryPg) UpdateUserByID(ctx context.Context, id string, user *entity.User) error {
	params := pg.UpdateUserByIDParams{
		ID:        StringToUUID(id),
		Email:     user.Email,
		Avatar:    pgtype.Text{String: string(user.Avatar), Valid: true},
		Username:  user.Username,
		Password:  user.Password,
		UpdatedAt: pgtype.Timestamp{Time: user.UpdatedAt, Valid: true},
	}

	return r.Queries.UpdateUserByID(ctx, params)
}

func (r *UserRepositoryPg) DeleteUserByID(ctx context.Context, id string) error {
	return r.Queries.DeleteUserByID(ctx, StringToUUID(id))
}
