package SQL

import (
	_ "database/sql"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"projects/LibraryDatabase2.0/TerminalClear"
	"projects/LibraryDatabase2.0/database"
)

type QuerySlice struct {
	QSId		string
	QSAuthor	string
	QSTitle		string
	QSPublish	string
	QSIsbn		string
}

var query QuerySlice

// QueryCatalog() calls for a listing of the SQL database onscreen,
// and creates a table via the go-pretty package.  Any calls for the
// display of the library will directly make a call to QueryCatalog()
func QueryCatalog() bool {
	TerminalClear.CallClear()
	rows, err := database.DBCon.Query("SELECT id, author, title, publish_date, isbn FROM library ORDER BY author ASC")
	if err != nil {
		fmt.Println("there was an error accessing the catalog: ", err)
	}
	t:= table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Author", "Title", "Published", "ISBN"})
	for rows.Next() {
		rows.Scan(&query.QSId, &query.QSAuthor, &query.QSTitle, &query.QSPublish, &query.QSIsbn)
		t.AppendRows([]table.Row{{query.QSId, text.WrapSoft(query.QSAuthor, 30), text.WrapSoft(query.QSTitle, 40), query.QSPublish, query.QSIsbn}})
	}
	t.SetColumnConfigs([]table.ColumnConfig{{Name: "Title", WidthMax: 40}})
	t.SetStyle(table.StyleColoredBlackOnGreenWhite)
	t.SetIndexColumn(1)
	t.SetTitle("Catalog of Books")
	t.AppendFooter(table.Row{"", "Library Total:", Count(), ""})
	t.Render()
	return true
}

// Count() provides the user with the total number of books that are
// contained in the database.
func Count() int {
	var count int
	row := database.DBCon.QueryRow("SELECT COUNT(*) FROM library")
	err := row.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}