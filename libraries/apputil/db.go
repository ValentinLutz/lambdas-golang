package apputil

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func NewDatabaseConfig(secret Secret) (*DatabaseConfig, error) {
	if secret.Username == "" {
		return nil, fmt.Errorf("secret username not set")
	}
	if secret.Password == "" {
		return nil, fmt.Errorf("secret password not set")
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, fmt.Errorf("env DB_HOST not set")
	}
	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, fmt.Errorf("env DB_PORT not set")
	}
	name, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, fmt.Errorf("env DB_NAME not set")
	}

	return &DatabaseConfig{
		Host:     host,
		Port:     port,
		Name:     name,
		User:     secret.Username,
		Password: secret.Password,
	}, nil
}

func NewDatabase(config *DatabaseConfig) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		config.Host, config.Port, config.Name, config.User, config.Password,
	)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
