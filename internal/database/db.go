package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "./agenda.db?_foreign_keys=on")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)

	initSchema(db)

	return db
}

func initSchema(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		phone TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}