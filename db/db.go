package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS data(id TEXT PRIMARY KEY, content TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func StoreData(id, content string) error {
	_, err := db.Exec("INSERT OR REPLACE INTO data(id, content) VALUES(?, ?)", id, content)
	return err
}

func GetData() (map[string]string, error) {
	rows, err := db.Query("SELECT id, content FROM data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := make(map[string]string)
	for rows.Next() {
		var id, content string
		if err := rows.Scan(&id, &content); err != nil {
			return nil, err
		}
		data[id] = content
	}
	return data, nil
}
