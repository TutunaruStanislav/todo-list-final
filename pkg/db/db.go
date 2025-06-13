package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var database *sql.DB

const schema string = `
	CREATE TABLE scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT "",
		title VARCHAR(100) NOT NULL,
		comment TEXT DEFAULT NULL,
		repeat VARCHAR(128) DEFAULT NULL
	);
	CREATE INDEX idx_scheduler_date ON scheduler (date);
`

func Init() error {
	var dbFile = "scheduler.db"
	path := os.Getenv("TODO_DBFILE")
	if len(path) > 0 {
		dbFile = path
	}
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	database, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if install {
		_, err := database.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}
