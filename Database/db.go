package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api_1.db")
	if err != nil {
		panic("Could not established the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	CreateTables()
}

func CreateTables() {
	CreateEventTables := `
		CREATE TABLE IF NOT EXISTS EVENT_TAB (
			id          INTEGER PRIMARY KEY AUTOINCREMENT,
			name        TEXT NOT NULL,
			description TEXT NOT NULL,
			location    TEXT NOT NULL,
			dateTime    DATETIME NOT NULL,
			userId      INTEGER
		)
	
	`
	_, err := DB.Exec(CreateEventTables)
	if err != nil {
		fmt.Println(err)
		panic("Cannot create tables")

	}

}
