package repository

import (
	"context"
	"github.com/firerplayer/stash-task/backend/internal/domain/entity"
	pg "github.com/firerplayer/stash-task/backend/internal/infra/pg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type TaskRepositoryPg struct {
	DB      *pgx.Conn
	Queries *pg.Queries
}

func NewTaskRepositoryPg(dbt *pgx.Conn) *TaskRepositoryPg {
	return &TaskRepositoryPg{
		DB:      dbt,
		Queries: pg.New(dbt),
	}
}

func (t TaskRepositoryPg) CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {

	taskDB, err := t.Queries.CreateTask(ctx, pg.CreateTaskParams{
		User:        StringToUUID(task.UserID),
		Title:       task.Title,
		Description: pgtype.Text{String: task.Description, Valid: true},
		Priority:    pgtype.Int4{Int32: int32(task.Priority), Valid: true},
		CompletedAt: pgtype.Timestamp{Time: task.CompletedAt, Valid: !task.CompletedAt.IsZero()},
		CreatedAt:   pgtype.Timestamp{Time: task.CreatedAt, Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: task.UpdatedAt, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	var taskEntity entity.Task
	err = HydrateTask(taskDB, &taskEntity)
	if err != nil {
		return nil, err
	}
	return &taskEntity, nil
}

func (t TaskRepositoryPg) GetTaskByID(ctx context.Context, id string) (*entity.Task, error) {
	taskDB, err := t.Queries.GetTaskByID(ctx, StringToUUID(id))
	if err != nil {
		return nil, err
	}
	var taskEntity entity.Task
	err = HydrateTask(taskDB, &taskEntity)
	if err != nil {
		return nil, err
	}
	return &taskEntity, nil
}

func (t TaskRepositoryPg) ListAllTasks(ctx context.Context, limit, offset int) ([]*entity.Task, error) {
	if limit <= 0 {
		limit = 20
	}
	allTasks, err := t.Queries.ListAllTasks(ctx, pg.ListAllTasksParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	var tasks []*entity.Task
	for _, taskDB := range allTasks {
		var task entity.Task
		err := HydrateTask(taskDB, &task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (t TaskRepositoryPg) ListTasksByUser(ctx context.Context, userID string, limit, offset int) ([]*entity.Task, error) {
	if limit <= 0 {
		limit = 20
	}
	tasksDB, err := t.Queries.ListTasksByUser(ctx, pg.ListTasksByUserParams{
		User:   StringToUUID(userID),
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	var tasks []*entity.Task
	for _, taskDB := range tasksDB {
		var task entity.Task
		err := HydrateTask(taskDB, &task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (t TaskRepositoryPg) CompleteTaskByID(ctx context.Context, id string) error {
	return t.Queries.CompleteTaskByID(ctx, pg.CompleteTaskByIDParams{
		CompletedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		ID:          StringToUUID(id),
	})

}

func (t TaskRepositoryPg) UncompleteTaskByID(ctx context.Context, id string) error {
	return t.Queries.UncompleteTaskByID(ctx, StringToUUID(id))

}

func (t TaskRepositoryPg) DeleteTaskByID(ctx context.Context, id string) error {
	return t.Queries.DeleteTaskByID(ctx, StringToUUID(id))

}

func (t TaskRepositoryPg) UpdateTaskByID(ctx context.Context, id string, task *entity.Task) error {
	params := pg.UpdateTaskByIDParams{
		Title:       task.Title,
		Description: pgtype.Text{String: task.Description, Valid: true},
		Priority:    pgtype.Int4{Int32: int32(task.Priority), Valid: true},
		CompletedAt: pgtype.Timestamp{Time: task.CompletedAt, Valid: true},
		UpdatedAt:   pgtype.Timestamp{Time: task.UpdatedAt, Valid: true},
		ID:          StringToUUID(id),
	}
	return t.Queries.UpdateTaskByID(ctx, params)
}
