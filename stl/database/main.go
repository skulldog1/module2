package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "./gopher.db")
	if err != nil {
		fmt.Println("Error preparing statement:", err.Error())
		return
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people(id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		fmt.Println("Error executing statement:", err.Error())
		return
	}
	statement.Exec()
	statement, err = database.Prepare("INSERT INTO people (first, lastname) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Error executing statement:", err.Error())
		return
	}
	statement.Exec("Lorem", "Ipsum")
	rows, err := database.Query("SELECT id, firstname, lastname FROM people")
	if err != nil {
		fmt.Println("Error executing statement:", err.Error())
		return
	}
	var id int
	var firstname string
	var lastname string

	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Printf("%d: %s %s\n", id, firstname, lastname)
	}
}
