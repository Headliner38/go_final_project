package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

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
const defaultPath = "scheduler.db"

// /Users/macbook/Desktop/go_final_project/scheduler.db
func dbTrack() string {
	pathDB := os.Getenv("TODO_DBFILE")
	if len(pathDB) == 0 {
		fmt.Printf("путь к БД из переменной окружения не задан, используется путь по умолчанию\n")

		workDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("ошибка получения рабочей директории")
			return defaultPath
		}

		return filepath.Join(workDir, defaultPath)
	}

	return pathDB
}

func Init() (*sql.DB, error) {
	path := dbTrack()
	_, err := os.Stat(path)
	exists := os.IsNotExist(err)

	db, err := sql.Open("sqlite3", path)
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
