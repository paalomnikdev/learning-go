package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Can't connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	tables := map[string]string{
		"users": `
			CREATE TABLE IF NOT EXISTS users (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				email TEXT NOT NULL UNIQUE,
				password TEXT NOT NULL
			)
		`,
		"events": `
			CREATE TABLE IF NOT EXISTS events (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL,
				description TEXT NOT NULL,
				location TEXT NOT NULL,
				dateTime DATETIME NOT NULL,
				user_id INTEGER,
				FOREIGN KEY(user_id) REFERENCES users(id)
			)
		`,
	}

	
	for table, query := range tables {
		_, err := DB.Exec(query)

		if err != nil {
			panic(
				fmt.Sprintf("Can't create %s table.", table),
			)
		}
	}
}
