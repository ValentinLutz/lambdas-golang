package testutil

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase(config *Config) *sqlx.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Database,
	)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func LoadAndExec(db *sqlx.DB, path string) {
	query := LoadQuery(path)
	Exec(db, query)
}

func LoadQuery(path string) string {
	query, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(query)
}

func Exec(db *sqlx.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
