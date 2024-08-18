package sqlite

import (
	"TODO_App/todo"
	"database/sql"
	"fmt"
	"time"

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
		event TEXT NOT NULL,
		day TEXT NOT NULL,
		time TEXT NOT NULL);
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

	res, err := s.db.Exec(`INSERT INTO todo(event, day, time) VALUES($1, $2, $3)`, event.Description, event.Date.Format("2006-01-02"), event.Date.Format("15:04"))
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return res.LastInsertId()
}

func (s *Storage) GetTodayTODOS() {
	today := time.Now().Format("2006-01-02")

	rows, err := s.db.Query(`SELECT event, time FROM todo WHERE day = $1 ORDER BY time;`, today)
	if err != nil {
		fmt.Println("Some err in getting events")
		return
	}
	for rows.Next() {
		var event string
		var time string

		err = rows.Scan(&event, &time)
		if err != nil {
			break
		}
		fmt.Printf("%s %s\n", event, time)
	}

}
