package sqlite

import (
	"TODO_App/todo"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(StoragePath string) (*Storage, error) {
	const op = "Storage.sqlite.new"

	db, err := sql.Open("sqlite3", StoragePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS todo(
		id INTEGER PRIMARY KEY,
		event TEXT NOT NULL UNIQUE,
		date TEXT NOT NULL);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) AddTODO(event todo.TODO) (int64, error) {
	const op = "storage.sqlite.AddTODO"

	date := strings.Join([]string{event.Date, event.Time}, "-")

	res, err := s.db.Exec(`INSERT INTO todo(event, date) VALUES($1, $2)`, event.Description, date)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return res.LastInsertId()
}
