package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func init() {
	// open db
	var err error
	Db, err = sql.Open("sqlite3", "./public/db/db.db")

	// handle error
	if err != nil {
		fmt.Println("Error opening the database:", err.Error())
		os.Exit(1)
	}
}
