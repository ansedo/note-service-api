package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Query struct {
	Name     string
	QueryRaw string
}

type DB struct {
	pool *pgxpool.Pool
}

func (db *DB) Get(ctx context.Context, dst any, q Query, args ...any) error {
	return db.pool.QueryRow(ctx, q.QueryRaw, args).Scan(dst)
}

func (db *DB) Select(ctx context.Context, dst any, q Query, args ...any) error {
	rows, err := db.pool.Query(ctx, q.QueryRaw, args)
	if err != nil {
		return err
	}
	return rows.Scan(dst)
}

func (db *DB) Exec(ctx context.Context, q Query, args ...any) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, q.QueryRaw, args)
}

func (db *DB) Query(ctx context.Context, q Query, args ...any) (pgx.Rows, error) {
	return db.pool.Query(ctx, q.QueryRaw, args)
}

func (db *DB) QueryRow(ctx context.Context, q Query, args ...any) pgx.Row {
	return db.pool.QueryRow(ctx, q.QueryRaw, args)
}
