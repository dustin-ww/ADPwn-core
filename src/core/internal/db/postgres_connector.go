package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
)

var (
	pgOnce sync.Once
	pgDB   *sql.DB
	pgErr  error
)

func GetPostgresDB() (*sql.DB, error) {
	pgOnce.Do(func() {
		connString := "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"

		db, err := sql.Open("pgx", connString)
		if err != nil {
			pgErr = fmt.Errorf("postgres connection failed: %w", err)
			return
		}

		// Test
		if err := db.PingContext(context.Background()); err != nil {
			pgErr = fmt.Errorf("postgres ping failed: %w", err)
			db.Close()
			return
		}

		pgDB = db
	})

	return pgDB, pgErr
}

func ExecutePostgresInTransaction(ctx context.Context, db *sql.DB, op func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("transaction start failed: %w", err)
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {

		}
	}(tx)

	if err := op(tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}
	return nil
}

func ExecutePostgresRead[T any](ctx context.Context, db *sql.DB, op func(tx *sql.Tx) (T, error)) (T, error) {
	var zero T
	tx, err := db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return zero, fmt.Errorf("read transaction start failed: %w", err)
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {

		}
	}(tx)

	result, err := op(tx)
	if err != nil {
		return zero, err
	}

	if err := tx.Commit(); err != nil {
		return zero, fmt.Errorf("read transaction commit failed: %w", err)
	}

	return result, nil
}
