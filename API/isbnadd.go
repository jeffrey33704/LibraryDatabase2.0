package API

import (
	"fmt"
	"projects/LibraryDatabase2.0/TerminalClear"
)

// User enters an ISBN10 or ISBN13.
func GetIsbn() string {
	TerminalClear.CallClear()
	var ISBN string
	fmt.Println("Please enter an ISBN13 or ISBN10: ")
	fmt.Scanln(&ISBN)
	return ISBN
}

// Checks to make sure the ISBN entered is valid.
func IsbnValid(ISBN string) (string, bool) {
	if len(ISBN) == 10 || len(ISBN) == 13 {
		return ISBN, true
	} else {
		return "", false
	}
}

// If the ISBN is valid, the ISBN number is looked up via
// googleapis. If the call returns a false, meaning googleapis
// hasn't a record of the ISBN, then a call is made to the
// openlibrary api.
// After receiving the answer to the request, JSON is parsed, and
// the relevant variables are sent back for processing into the
// SQLite database.
func AddISBN(isbn string) (string, string, string, string) {
	var DataReceived bool
	var Title, Author, PubDate, ISBN13 string
	_, _, _, _, DataReceived = GoogleAPI(isbn)
	if DataReceived == false {
		Title, Author, PubDate, ISBN13, _ = OpenLibrary(isbn)
		fmt.Println("Error finding ISBN")
	} else {
		Title, Author, PubDate, ISBN13, _ = GoogleAPI(isbn)
	}

	return Title, Author, PubDate, ISBN13
}