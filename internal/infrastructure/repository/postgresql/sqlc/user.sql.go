// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const findUser = `-- name: FindUser :one
SELECT id, email, created_at FROM users WHERE id = $1
`

func (q *Queries) FindUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, findUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.CreatedAt)
	return i, err
}

const saveUser = `-- name: SaveUser :one
INSERT INTO users (id, email, created_at) VALUES ($1, $2, $3) RETURNING id, email, created_at
`

type SaveUserParams struct {
	ID        uuid.UUID
	Email     string
	CreatedAt pgtype.Timestamp
}

func (q *Queries) SaveUser(ctx context.Context, arg SaveUserParams) (User, error) {
	row := q.db.QueryRow(ctx, saveUser, arg.ID, arg.Email, arg.CreatedAt)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.CreatedAt)
	return i, err
}
