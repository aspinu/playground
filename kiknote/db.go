package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DataStore struct {
	db *sql.DB
}

type Note struct {
	Id    int
	Title string
	Body  string
}

func (store *DataStore) Init() error {
	var err error
	store.db, err = sql.Open("sqlite3", "./notes.db")
	if err != nil {
		return err
	}
	// defer store.db.Close()
	stetment := `CREATE TABLE IF NOT EXISTS notes (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT not null,
		body TEXT not null
	);`
	if _, err := store.db.Exec(stetment); err != nil {
		return err
	}
	return nil
}

func (store *DataStore) GetNotes() ([]Note, error) {
	rows, err := store.db.Query("SELECT * FROM notes ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := make([]Note, 0)
	for rows.Next() {
		note := Note{}
		err = rows.Scan(&note.Id, &note.Title, &note.Body)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)

	}
	return notes, nil
}

func (store *DataStore) SaveNotes(note Note) error {
	// curentTime := time.Now().Format("2017-09-07")
	if note.Id == 0 {
		// pseudo-unique id
		note.Id = time.Now().Minute()
	}
	stmt, err := store.db.Prepare("INSERT INTO notes (id, title, body ) VALUES (?, ?, ? ) ON CONFLICT(id) DO UPDATE SET title=excluded.title, body=excluded.body;")
	if err != nil {
		return err
	}
	stmt.Exec(nil, note.Title, note.Body)
	defer stmt.Close()

	return nil
}
