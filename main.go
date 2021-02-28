package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"projects/LibraryDatabase2.0/TerminalClear"
	"projects/LibraryDatabase2.0/database"
	"projects/LibraryDatabase2.0/menu"
)

// Establishes connection to SQLite database, keeps connection open,
// and calls forth the Menu().
func main() {
	var err error
	database.DBCon, err = sql.Open("sqlite3", "./library.db")
	if err != nil {
		fmt.Println("there was an error: ", err)
	}
	defer database.DBCon.Close()
	statement, _ := database.DBCon.Prepare("CREATE TABLE IF NOT EXISTS library (id INTEGER PRIMARY KEY, author TEXT, title TEXT, publish_date TEXT, isbn TEXT)")
	statement.Exec()
	TerminalClear.CallClear()
	menu.Menu()
}