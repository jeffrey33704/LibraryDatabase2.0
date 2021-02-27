package SQL

import (
	"fmt"
	"projects/LibraryDatabase2.0/TerminalClear"
	"projects/LibraryDatabase2.0/database"
	"strings"
)

// Variables are inputted from ISBNAdd() and are entered in the database.
func AddBook(Title, Author, PubDate, ISBN string) {
	if strings.TrimSpace(Title) == "" {
		fmt.Println("Unable to add to library")
		return
	}
	statement, _ := database.DBCon.Prepare("CREATE TABLE IF NOT EXISTS library (id INTEGER PRIMARY KEY, author TEXT, title TEXT, publish_date TEXT, isbn TEXT)")
	statement.Exec()
	statement, _ = database.DBCon.Prepare("INSERT INTO library (author, title, publish_date, isbn) VALUES (?, ?, ?, ?)")
	statement.Exec(Author, Title, PubDate, ISBN)

	TerminalClear.CallClear()
	QueryCatalog()
}

// If the user wishes to delete an entry, this function does so.
func DeleteBook() {
	QueryCatalog()
	var entry int
	fmt.Println("Please enter the number to delete: ")
	fmt.Scanln(&entry)
	_, err := database.DBCon.Exec("DELETE FROM library WHERE id = ?", entry)
	if err != nil {
		fmt.Println(err)
	}
	TerminalClear.CallClear()
	QueryCatalog()
}