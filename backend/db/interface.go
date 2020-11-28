package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Queryable interface {
	QueryRow(ctx context.Context, sql string, params ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, params ...interface{}) (pgconn.CommandTag, error)
}

type Transactable interface {
	Queryable
	Begin(ctx context.Context) (pgx.Tx, error)
}
