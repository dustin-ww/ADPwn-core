package main

import (
	"ADPwn-core/internal/db"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"os"
	"strings"
	"time"
)

func InitializePostgresDB() {
	gormDB, err := db.GetPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get the underlying *sql.DB from GORM
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB from GORM: %v", err)
	}

	// Now you can use *sql.DB methods
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	schema, err := os.ReadFile("./adpwn.sql")
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}
	log.Println("Postgres: Schema file read successfully.")

	commands := strings.Split(string(schema), ";")
	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}

		result := gormDB.Exec(cmd)
		if result.Error != nil {
			log.Fatalf("Failed to execute SQL command: %v", result.Error)
		}
	}
	log.Println("Postgres: Database schema initialized successfully")
}

func DropAllPostgresObjects() error {
	gormDB, err := db.GetPostgresDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// The SQL script for dropping everything
	dropScript := `
    -- Disable foreign key checks temporarily
    SET session_replication_role = 'replica';
    
    -- Drop all tables
    DO $$ 
    DECLARE
        r RECORD;
    BEGIN
        FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
            EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
        END LOOP;
    END $$;
    
    -- Drop all views
    DO $$ 
    DECLARE
        r RECORD;
    BEGIN
        FOR r IN (SELECT viewname FROM pg_views WHERE schemaname = 'public') LOOP
            EXECUTE 'DROP VIEW IF EXISTS ' || quote_ident(r.viewname) || ' CASCADE';
        END LOOP;
    END $$;
    
    -- Drop all sequences
    DO $$ 
    DECLARE
        r RECORD;
    BEGIN
        FOR r IN (SELECT sequence_name FROM information_schema.sequences WHERE sequence_schema = 'public') LOOP
            EXECUTE 'DROP SEQUENCE IF EXISTS ' || quote_ident(r.sequence_name) || ' CASCADE';
        END LOOP;
    END $$;
    
    -- Drop all types
    DO $$ 
    DECLARE
        r RECORD;
    BEGIN
        FOR r IN (SELECT typname FROM pg_type WHERE typnamespace = (SELECT oid FROM pg_namespace WHERE nspname = 'public') AND typtype = 'c') LOOP
            EXECUTE 'DROP TYPE IF EXISTS ' || quote_ident(r.typname) || ' CASCADE';
        END LOOP;
    END $$;
    
    -- Reset foreign key checks
    SET session_replication_role = 'origin';
    `

	result := gormDB.Exec(dropScript)
	if result.Error != nil {
		return fmt.Errorf("failed to drop database objects: %w", result.Error)
	}

	log.Println("Postgres: All database objects dropped successfully")
	return nil
}
