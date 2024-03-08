// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package model

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.getFullNameByIdStmt, err = db.PrepareContext(ctx, getFullNameById); err != nil {
		return nil, fmt.Errorf("error preparing query GetFullNameById: %w", err)
	}
	if q.getUserForLoginStmt, err = db.PrepareContext(ctx, getUserForLogin); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserForLogin: %w", err)
	}
	if q.getUsersByQueryStmt, err = db.PrepareContext(ctx, getUsersByQuery); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsersByQuery: %w", err)
	}
	if q.listUserStmt, err = db.PrepareContext(ctx, listUser); err != nil {
		return nil, fmt.Errorf("error preparing query ListUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.getFullNameByIdStmt != nil {
		if cerr := q.getFullNameByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFullNameByIdStmt: %w", cerr)
		}
	}
	if q.getUserForLoginStmt != nil {
		if cerr := q.getUserForLoginStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserForLoginStmt: %w", cerr)
		}
	}
	if q.getUsersByQueryStmt != nil {
		if cerr := q.getUsersByQueryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersByQueryStmt: %w", cerr)
		}
	}
	if q.listUserStmt != nil {
		if cerr := q.listUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                  DBTX
	tx                  *sql.Tx
	createUserStmt      *sql.Stmt
	getFullNameByIdStmt *sql.Stmt
	getUserForLoginStmt *sql.Stmt
	getUsersByQueryStmt *sql.Stmt
	listUserStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                  tx,
		tx:                  tx,
		createUserStmt:      q.createUserStmt,
		getFullNameByIdStmt: q.getFullNameByIdStmt,
		getUserForLoginStmt: q.getUserForLoginStmt,
		getUsersByQueryStmt: q.getUsersByQueryStmt,
		listUserStmt:        q.listUserStmt,
	}
}
