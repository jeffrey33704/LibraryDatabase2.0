package API

import (
	"bufio"
	"fmt"
	fp "github.com/amonsat/fullname_parser"
	"os"
	"strings"
)

// If the user does not have access to an ISBN number, if the book was
// published prior to ISBN, or in the unlikely event that neither
// googleapis nor openlibrary has an entry, the user may add a book manually.
func ManualAdd() (string, string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the author's name: ")
	scanner.Scan()
	Author := scanner.Text()
	parsedFullName := fp.ParseFullname(Author)
	parsedAuthor := fmt.Sprintf(parsedFullName.Last + ", " + parsedFullName.First + " " + parsedFullName.Middle)
	Author = parsedAuthor

	fmt.Println("Is there an additional author? Y/N")
	scanner.Scan()
	additional := scanner.Text()
	if additional == "Y" || additional == "y" {
		fmt.Println("Please enter the second author's name: ")
		scanner.Scan()
		SecondAuthor := scanner.Text()
		Author += " & " + SecondAuthor
	}

	fmt.Println("Enter the title of the book: ")
	scanner.Scan()
	Title := scanner.Text()
	fmt.Println("Enter the subtitle, if any: ")
	scanner.Scan()
	Subtitle := scanner.Text()
	blank := strings.TrimSpace(Subtitle) == ""
	if blank == false {
		Title = Title + ": " + Subtitle
	}

	fmt.Println("Enter the year of publication: ")
	scanner.Scan()
	PubDate := scanner.Text()

	fmt.Println("Enter the ISBN number: ")
	scanner.Scan()
	ISBN := scanner.Text()

	return Title, Author, PubDate, ISBN
}