// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package model

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO user (
  username, password
) VALUES (
  ?, ?
)
RETURNING id, username
`

type CreateUserParams struct {
	Username string
	Password string
}

type CreateUserRow struct {
	ID       int64
	Username string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.Username, arg.Password)
	var i CreateUserRow
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, username FROM user
ORDER BY name
`

type ListUserRow struct {
	ID       int64
	Username string
}

func (q *Queries) ListUser(ctx context.Context) ([]ListUserRow, error) {
	rows, err := q.query(ctx, q.listUserStmt, listUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUserRow{}
	for rows.Next() {
		var i ListUserRow
		if err := rows.Scan(&i.ID, &i.Username); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}