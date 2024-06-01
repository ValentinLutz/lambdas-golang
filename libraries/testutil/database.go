package testutil

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
)

func MustLoadAndExec(db *pgxpool.Pool, path string) {
	query := MustLoadQuery(path)
	MustExec(db, query)
}

func MustLoadQuery(path string) string {
	query, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("failed to read file: %w", err))
	}

	return string(query)
}

func MustExec(db *pgxpool.Pool, query string) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()

	_, err := db.Exec(ctx, query)
	if err != nil {
		panic(fmt.Errorf("failed to execute query: %w", err))
	}
}
