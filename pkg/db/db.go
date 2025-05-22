package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '',
	title VARCHAR,
	comment TEXT,
	repeat VARCHAR
);

CREATE INDEX date_idx ON scheduler(date);
`

func Init(dbFile string) (*sql.DB, error) {
	_, err := os.Stat(dbFile)
	exists := os.IsNotExist(err)

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return db, err
	}
	//defer db.Close()
	if exists {
		_, err := db.Exec(schema)
		if err != nil {
			return db, err
		}
	}
	return db, nil
}
