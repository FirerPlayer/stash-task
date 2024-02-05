package repository

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/firerplayer/stash-task/backend/internal/domain/entity"
	pg "github.com/firerplayer/stash-task/backend/internal/infra/pg"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// HydrateTask hydrates a Task object from a database Task object
func HydrateTask(taskDB pg.Task, task *entity.Task) error {
	var err error
	task.ID, err = uuid.Parse(UUIDToString(taskDB.ID))
	if err != nil {
		return errors.New("failed to hydrate task --> " + err.Error())
	}
	task.UserID = UUIDToString(taskDB.User)
	task.Title = taskDB.Title
	task.Description = taskDB.Description.String
	task.Priority = int(taskDB.Priority.Int32)
	task.CompletedAt = taskDB.CompletedAt.Time
	task.CreatedAt = taskDB.CreatedAt.Time
	task.UpdatedAt = taskDB.UpdatedAt.Time
	return nil
}

// HydrateUser hydrates a user entity with data from a userDb record.
//
// userDb: The database record to hydrate from.
// user: The user entity to hydrate.
// error: Returns an error if the JSON unmarshal fails.
func HydrateUser(userDb pg.User, user *entity.User) error {
	var err error
	user.ID, err = uuid.Parse(UUIDToString(userDb.ID))
	if err != nil {
		return errors.New("failed to hydrate user --> " + err.Error())
	}
	user.Email = userDb.Email
	user.Avatar = []byte(userDb.Avatar.String)
	user.Username = userDb.Username
	user.Password = userDb.Password
	user.Bio = userDb.Bio.String
	user.CreatedAt = userDb.CreatedAt.Time
	user.UpdatedAt = userDb.UpdatedAt.Time
	return nil
}

func UUID(v uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: v,
		Valid: true,
	}
}

// StringToUUID converts a string UUID in standard form to a pgtype.UUID.
// Check for Valid before using the result.
func StringToUUID(s string) pgtype.UUID {
	data, err := parseUUID(s)
	if err != nil {
		return pgtype.UUID{
			Bytes: [16]byte{},
			Valid: false,
		}
	}
	return pgtype.UUID{
		Bytes: data,
		Valid: true,
	}
}

// parseUUID converts a string UUID in standard form to a byte array.
func parseUUID(src string) (dst [16]byte, err error) {
	switch len(src) {
	case 36:
		src = src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:]
	case 32:
		// dashes already stripped, assume valid
	default:
		// assume invalid.
		return dst, fmt.Errorf("cannot parse UUID %v", src)
	}

	buf, err := hex.DecodeString(src)
	if err != nil {
		return dst, err
	}

	copy(dst[:], buf)
	return dst, err
}

// UUIDToString returns format xxxx-yyyy-zzzz-rrrr-tttt
func UUIDToString(t pgtype.UUID) string {
	if !t.Valid {
		return ""
	}
	src := t.Bytes
	return fmt.Sprintf("%x-%x-%x-%x-%x", src[0:4], src[4:6], src[6:8], src[8:10], src[10:16])
}
