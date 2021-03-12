package menu

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"projects/LibraryDatabase2.0/API"
	"projects/LibraryDatabase2.0/SQL"
	"projects/LibraryDatabase2.0/TerminalClear"
	"time"
)

var Title, Author, PubDate, ISBN13 string

// The "Hub" of the program.  Menu() allows the user to select from
// a list of options to manage their library database.
func Menu() {
	prompt := promptui.Select{
		Label: "Select and option: ",
		Items: []string{
			"List Catalog",
			"Add New (ISBN)",
			"Add New (Manual)",
			"Scan ISBN",
			"Delete Entry",
			"Quit",
		},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	switch result {
	case "List Catalog":
		true := SQL.QueryCatalog()
		if true {Menu()}
	case "Add New (ISBN)":
		getISBN := API.GetIsbn()
		validISBN, status := API.IsbnValid(getISBN)
		if status == false {
			fmt.Println("invalid ISBN entered")
			fmt.Println("returning to main menu")
			time.Sleep(3 * time.Second)
			TerminalClear.CallClear()
			Menu()
		}
		if status == true {
			SQL.AddBook(API.AddISBN(validISBN))
		}
		Menu()
	case "Add New (Manual)":
		TerminalClear.CallClear()
		SQL.AddBook(API.ManualAdd())
		Menu()
	case "Scan ISBN":
		TerminalClear.CallClear()
		WebcamISBN := API.ExportISBN()
		SQL.AddBook(API.AddISBN(WebcamISBN))
		Menu()
	case "Delete Entry":
		SQL.DeleteBook()
		Menu()
	case "Quit":
		TerminalClear.CallClear()
		os.Exit(0)
	}
}