package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var database *sql.DB

// schema describes the structure of the DB
const schema string = `
	CREATE TABLE scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT "",
		title VARCHAR(100) NOT NULL COLLATE NOCASE,
		comment TEXT DEFAULT NULL COLLATE NOCASE,
		repeat VARCHAR(128) DEFAULT NULL COLLATE NOCASE
	);
	CREATE INDEX idx_scheduler_date ON scheduler (date);
`

// initDataDir - create data/ directory in the root of the project to store data.
func initDataDir() error {
	dirPath := "./data"
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// Init returns a pointer to DB if successful, otherwise DB initialization error.

// Before running you must define the path to the SQlite DB file
// in the TODO_DBFILE environment variable, the function will create it automatically
// if it doesn't exist, otherwise it will initialize a connection to it.
func Init() (*sql.DB, error) {
	var install bool

	err := initDataDir()
	if err != nil {
		return nil, err
	}

	path := os.Getenv("TODO_DBFILE")
	_, err = os.Stat(path)
	if err != nil {
		install = true
	}

	database, err = sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if install {
		_, err := database.Exec(schema)
		if err != nil {
			return nil, err
		}
	}

	return database, nil
}
