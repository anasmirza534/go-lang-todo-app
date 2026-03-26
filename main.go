package main

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	const sqliteFile string = "./todo.db"
	db, err := sql.Open("sqlite", sqliteFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Database connected")

	row := db.QueryRow("SELECT 25;")
	var result int
	if err := row.Scan(&result); err != nil {
		log.Fatal(err)
	}

	log.Println("Query result", result)
}

// current goal:
// ☑️ sqlite setup
// ☑️ data modeling
//    migration setup
//    db query
//    web server setup
//    health route setup
//    rest api
