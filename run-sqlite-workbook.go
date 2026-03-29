package main

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

func RunSqliteWorkbook() {
	const sqliteFile string = "./todo.db"
	db, err := sql.Open("sqlite", sqliteFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	log.Println("Database connected")

	row := db.QueryRow("SELECT 25;")
	var result int
	if err := row.Scan(&result); err != nil {
		log.Fatal(err)
	}

	log.Println("Query result", result)

	todos, err := ListAllTodos(db)
	if err != nil {
		log.Fatal(err)
	}

	todo, err := AddTodo(db, "buy groceries")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted todo: ", todo)

	if len(todos) == 0 {
		log.Println("No todo found in db.")
	}

	for idx, todo := range todos {
		log.Println(idx, todo)
	}

	// err = DeleteTodo(db, "ef860e72-cac6-4cb5-949c-5e0085f0e599")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	todo, err = GetTodo(db, "7a141931-1ed7-4f72-b612-07c63191de49")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(todo)
	err = ToggleTodo(db, "7a141931-1ed7-4f72-b612-07c63191de49")
	if err != nil {
		log.Fatal(err)
	}
	todo, err = GetTodo(db, "7a141931-1ed7-4f72-b612-07c63191de49")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(todo)
}

func ListAllTodos(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, done, created_at FROM todo;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []Todo{}
	for rows.Next() {
		todo, err := scanTodo(rows)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func AddTodo(db *sql.DB, title string) (Todo, error) {
	id := uuid.New().String()
	result, err := db.Exec("INSERT INTO todo (id, title) VALUES (?, ?)", id, title)
	if err != nil {
		return Todo{}, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return Todo{}, err
	}

	log.Println("Inserted records: ", count)

	return GetTodo(db, id)
}

func GetTodo(db *sql.DB, id string) (Todo, error) {
	row := db.QueryRow("SELECT id, title, done, created_at FROM todo WHERE id = ?;", id)
	todo, err := scanTodo(row)
	if err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func DeleteTodo(db *sql.DB, id string) error {
	result, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Println("Deleted records: ", count)
	if count != 1 {
		return errors.New("No record found.")
	}

	return nil
}

func ToggleTodo(db *sql.DB, id string) error {
	result, err := db.Exec("UPDATE todo SET done = NOT done WHERE id = ?", id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Println("Toggle todo records: ", count)
	if count != 1 {
		return errors.New("No record found.")
	}

	return nil
}

type RowScanner interface {
	Scan(dest ...any) error
}

func scanTodo(s RowScanner) (Todo, error) {
	var t Todo
	var createdAtStr string
	var err error

	if err = s.Scan(&t.ID, &t.Title, &t.Done, &createdAtStr); err != nil {
		return Todo{}, err
	}

	t.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return Todo{}, err
	}

	return t, nil
}

// current goal:
// ☑️ sqlite setup
// ☑️ data modeling
// ☑️ migration setup
// ☑️ db query
//    cli setup
//
// phase 2
//    web server setup
//    health route setup
//    rest api
