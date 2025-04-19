package db

import (
	setup "ADPwn-core/init"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	pgOnce sync.Once
	pgDB   *gorm.DB
	pgErr  error
)

type ctxKey string

const txKey ctxKey = "dbTx"

func GetPostgresDB() (*gorm.DB, error) {
	pgOnce.Do(func() {
		// Connection String
		host := os.Getenv("POSTGRES_HOST")
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")
		port := os.Getenv("POSTGRES_PORT")

		if host == "" {
			host = "localhost"
			log.Println("WARNING: POSTGRES_HOST not set, using localhost")
		}

		if user == "" || password == "" || dbname == "" || port == "" {
			user = "adpwn"
			password = "adpwn"
			dbname = "adpwn"
			port = "5432"
			log.Println("WARNING: POSTGRES vars via ens not set, trying postgres defaults")
		}

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, password, dbname, port,
		)

		config := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}

		db, err := gorm.Open(postgres.Open(dsn), config)

		if err != nil {
			pgErr = fmt.Errorf("gorm connection failed: %w", err)
			return
		}

		if !isPotgresInitialized(db) {
			log.Println("WARNING: POSTGRES is not initialized yet! Initializing database...")
			err := setup.InitializePostgresSchema(db)
			if err != nil {
				pgErr = fmt.Errorf("gorm connection failed: %w", err)
				return
			}
		}

		sqlDB, err := db.DB()
		if err != nil {
			pgErr = fmt.Errorf("connection pool setup failed: %w", err)
			return
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		pgDB = db
	})

	return pgDB, pgErr
}

func isPotgresInitialized(db *gorm.DB) bool {
	var count int64

	db.Raw("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public'").Scan(&count)

	return count > 0
}

func ExecutePostgresInTransaction(ctx context.Context, db *gorm.DB, op func(tx *gorm.DB) error) error {
	return db.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, txKey, tx)
		return op(tx.WithContext(ctx))
	})
}

func ExecutePostgresRead[T any](ctx context.Context, db *gorm.DB, op func(tx *gorm.DB) (T, error)) (T, error) {
	var result T

	// ReadOnly Transaction
	tx := db.Session(&gorm.Session{
		Context:     ctx,
		PrepareStmt: false,
	})

	err := tx.Transaction(func(tx *gorm.DB) error {
		tmp, err := op(tx)
		if err != nil {
			return err
		}
		result = tmp
		return nil
	})

	return result, err
}

func GetTxFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok {
		return tx
	}
	return pgDB.WithContext(ctx)
}
