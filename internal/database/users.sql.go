// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUsers = `-- name: createUsers :one
INSERT INTO users (ID, created_at, updated_at, name)
VALUES($1, $2, $3, $4)
RETURNING id, created_at, updated_at, name
`

type createUsersParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (q *Queries) createUsers(ctx context.Context, arg createUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}
