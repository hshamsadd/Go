package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq" // Postgres driver
)

func Connect(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	// Modern Best Practice: Connection Pooling
	// This prevents your app from crashing under heavy load
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify the connection is actually working
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
