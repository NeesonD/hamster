package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDb() {
	dbt, err := sql.Open("sqlite3", "./hamster.db")
	if err != nil {
		log.Fatal(err)
	}
	db = dbt
}
