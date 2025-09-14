package db

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnDBStr() string {
	dbUser := strings.Trim(os.Getenv("DB_USER"), "\"=")
	dbPassword := strings.Trim(os.Getenv("DB_PASSWORD"), "\"=")
	dbHost := strings.Trim(os.Getenv("DB_HOST"), "\"=")
	dbPort := strings.Trim(os.Getenv("DB_PORT"), "\"=")
	dbName := strings.Trim(os.Getenv("DB_NAME"), "\"=")
	dbSSLMode := strings.Trim(os.Getenv("DB_SSL_MODE"), "\"=")

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		dbSSLMode,
	)

	return connStr
}

func NewDBConfig(connStr string) (*pgxpool.Config, error) {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 5 * time.Minute

	return config, nil
}

func NewDBPool(config *pgxpool.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %v", err)
	}

	return pool, nil
}
