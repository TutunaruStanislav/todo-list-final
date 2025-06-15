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
