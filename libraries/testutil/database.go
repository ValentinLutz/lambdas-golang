package testutil

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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
