package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(ctx context.Context, driverName, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, driverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("sqlx.ConnectContext(): %w", err)
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db.PingContext(): %w", err)
	}

	return db, nil
}
