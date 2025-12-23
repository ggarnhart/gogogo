package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context) (*DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	config.MaxConns = 10
	config.MinConns = 2

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// Test Connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &DB{Pool: pool}, nil
}

func (db *DB) Close(ctx context.Context) error {
	db.Pool.Close()
	return nil
}
