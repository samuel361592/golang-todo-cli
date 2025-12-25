package todo

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const dbPath = "data/todos.db"

func ensureDataDir() error {
	return os.MkdirAll("data", 0755)
}

func openDB() (*sql.DB, error) {
	if err := ensureDataDir(); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed INTEGER NOT NULL DEFAULT 0
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// 覆蓋存檔：先清空再寫入（最貼近你原本 JSON 行為）
func SaveTodos(todos []Todo) error {
	db, err := openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec("DELETE FROM todos"); err != nil {
		tx.Rollback()
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO todos(title, completed) VALUES(?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, t := range todos {
		completed := 0
		if t.Completed {
			completed = 1
		}
		if _, err := stmt.Exec(t.Title, completed); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func LoadTodos() ([]Todo, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT title, completed FROM todos ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var title string
		var completed int
		if err := rows.Scan(&title, &completed); err != nil {
			return nil, err
		}
		todos = append(todos, Todo{
			Title:     title,
			Completed: completed == 1,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
