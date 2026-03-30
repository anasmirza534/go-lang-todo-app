package store

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Connect() (*sql.DB, error) {
	const sqliteFile string = "./todo.db"
	db, err := sql.Open("sqlite", sqliteFile)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	log.Println("Database connected")

	return db, nil
}
