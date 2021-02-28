package API

import (
	"encoding/json"
	"fmt"
	"github.com/amonsat/fullname_parser"
	"github.com/araddon/dateparse"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var GTitle, GAuthor, GPubDate, GISBN string

type Result struct {
	TotalItems	int	`json:"totalItems"`
	Items	[]struct {
		VolumeInfo struct {
			Title			string		`json:"title"`
			Subtitle		string		`json:"subtitle,omitempty"`
			Authors			[]string	`json:"authors,omitempty"`
			PublishedDate	string		`json:"publishedDate"`
			IndustryIdentifiers	[]struct {
				Type		string	`json:"type,omitempty"`
				Identifier	string	`json:"identifier,omitempty"`
			} `json:"industryIdentifiers,omitempty"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

type TotalItems struct {
	TotalResults	int
}

type Book struct {
	Title		string
	Subtitle	string
	Authors		[]string
	PubDate		string
	ISBN		string
}

// An ISBN is sent to the GoogleAPI function, which calls a "GET"
// request with the said ISBN. If the request times out, or if
// googleapis cannot locate the information associated with the
// ISBN, the request is sent back to ISBNAdd() for additional processing.
// If the request is successful, the relevant variables are sent back
// to ISBNAdd() for transfer into the SQLite database.
func GoogleAPI(isbn string) (string, string, string, string, bool) {
	lookup := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s", isbn)
	myClient := http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest(http.MethodGet, lookup, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Test API")
	res, getErr := myClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	result := Result{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	totalResults := TotalItems{
		TotalResults: result.TotalItems,
	}

	if totalResults.TotalResults < 1 {
		return "", "", "", "", false
	}

	if len(result.Items[0].VolumeInfo.Authors) == 0 {
		return "", "", "", "", false
	}

	book := Book{
		Title: result.Items[0].VolumeInfo.Title,
		Subtitle: result.Items[0].VolumeInfo.Subtitle,
		Authors: result.Items[0].VolumeInfo.Authors,
		PubDate: result.Items[0].VolumeInfo.PublishedDate,
	}

	parsedAuthor := fullname_parser.ParseFullname(book.Authors[0])
	parsed := fmt.Sprintf(parsedAuthor.Last + ", " + parsedAuthor.First + " " + parsedAuthor.Middle)
	if len(book.Authors) == 1 {
		GAuthor = parsed
	} else {
		GAuthor = parsed + " & " + book.Authors[1]
	}
	blank := strings.TrimSpace(book.Subtitle) == ""
	if blank == false {
		GTitle = book.Title + ": " + book.Subtitle
	} else {
		GTitle = book.Title
	}
	if len(book.PubDate) > 4 {
		d, _ := dateparse.ParseLocal(book.PubDate)
		GPubDate = d.Format("2006")
	} else {
		GPubDate = book.PubDate
	}

	if result.Items[0].VolumeInfo.IndustryIdentifiers[0].Type == "ISBN_13" {
		GISBN = result.Items[0].VolumeInfo.IndustryIdentifiers[0].Identifier
	} else if result.Items[0].VolumeInfo.IndustryIdentifiers[0].Type == "ISBN_10" {
		GISBN = result.Items[0].VolumeInfo.IndustryIdentifiers[1].Identifier
	} else {
		GISBN = isbn
	}

	return GTitle, GAuthor, GPubDate, GISBN, true
}