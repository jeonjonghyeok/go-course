package db

import (
	"database/sql"

	_ "github.com/sql/pq"
)

var DB *sql.DB

func Connect(url string) error {
	c, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}
	DB = c
	return nil
}
