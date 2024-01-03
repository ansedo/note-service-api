package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Client struct {
	db        *DB
	closeFunc context.CancelFunc
}

func NewClient(ctx context.Context, config *pgxpool.Config) (*Client, error) {
	dbc, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	_, cancel := context.WithCancel(ctx)

	return &Client{
		db:        &DB{pool: dbc},
		closeFunc: cancel,
	}, nil
}

func (c *Client) DB() *DB {
	return c.db
}

func (c *Client) Close() error {
	if c != nil {
		if c.closeFunc != nil {
			c.closeFunc()
		}

		if c.db != nil {
			c.db.pool.Close()
		}
	}

	return nil
}
