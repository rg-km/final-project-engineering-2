package main

import (
	"database/sql"
	"final-project-eng2-be/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./beasiswa.db")
	if err != nil {
		panic(err)
	}
	database.Migrate(db)

}
