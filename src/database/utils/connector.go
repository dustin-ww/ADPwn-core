package utils

import (
	"context"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db    *sqlx.DB
	once  sync.Once
	dbErr error
)

func GetDB() (*sqlx.DB, error) {
	once.Do(func() {
		db, dbErr = sqlx.Connect("sqlite3", "./../../db.sqlite")
		if dbErr != nil {
			log.Printf("Error while connecting to the database: %v", dbErr)
		}
	})
	return db, dbErr
}

func CloseDB() {
	if db != nil {
		_ = db.Close()
	}
}

func QueryContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	if db == nil {
		return nil, dbErr
	}
	return db.QueryxContext(ctx, query, args...)
}
