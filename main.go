package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	checkError(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	checkError(err)
	defer rows.Close()

	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		log.Println(id, name)
	}
}
