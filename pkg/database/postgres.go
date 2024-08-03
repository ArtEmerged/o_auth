package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPostgres creates a new postgres connection.
func NewPostgres(ctx context.Context, dbDNS string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(ctx, dbDNS)

	if err != nil {
		return nil, fmt.Errorf("failed connection to postgres db: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed ping to postgres db: %w", err)
	}

	return conn, nil
}
