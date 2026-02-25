package db

import (
	"bitalikr1999/difare/app/db/migrations"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLite(ctx context.Context, path string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s?_journal_mode=WAL&_busy_timeout=5000&_foreign_keys=on", path)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	migrations.MigrateTaskTable(db)

	fmt.Println("asdaswd")

	return db, nil
}
