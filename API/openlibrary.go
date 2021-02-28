package API

import (
	"encoding/json"
	"fmt"
	fp "github.com/amonsat/fullname_parser"
	"github.com/araddon/dateparse"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type OLData struct {
	Title	string	`json:"title"`
	Subtitle	string	`json:"subtitle,omitempty"`
	Authors	Authors
	Identifiers	Identifiers
	PubDate	string	`json:"publish_date"`
}

type Authors []struct {
	Name	string	`json:"name,omitempty"`
}

type Identifiers struct {
	ISBN10	[]string	`json:"isbn_10,omitempty"`
	ISBN13	[]string	`json:"isbn_13,omitempty"`
}


var OLTitle, OLSubtitle, OLAuthor, OLPubDate, OLISBN string

// If the GET request from GoogleAPI() fails, a GET request is made
// through OpenLibrary(). The requested information is then sent back
// to ISBNAdd() for processing into the SQLite database.
func OpenLibrary(isbn string) (string, string, string, string, bool) {
	lookup := fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%s&jscmd=data&format=json", isbn)
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

	data := map[string]OLData{}
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	for _, v := range data {
		OLTitle = v.Title
		OLSubtitle = v.Subtitle
		blank := strings.TrimSpace(v.Subtitle) == ""
		if blank == false {
			OLTitle = OLTitle + ": " + OLSubtitle
		}
		OLAuthor = v.Authors[0].Name
		ParsedFullName := fp.ParseFullname(OLAuthor)
		OLAuthor = fmt.Sprintf(ParsedFullName.Last + ", " + ParsedFullName.First + " " + ParsedFullName.Middle)
		if len(v.Authors) > 1 {
			Author2 := v.Authors[1].Name
			OLAuthor = OLAuthor + " & " + Author2
		}
		OLPubDate = v.PubDate
		if len(OLPubDate) > 4 {
			d, _ := dateparse.ParseLocal(OLPubDate)
			OLPubDate = d.Format("2006")
		}

		if len(v.Identifiers.ISBN10) == 0 {
			OLISBN = v.Identifiers.ISBN13[0]
		} else {
			OLISBN = v.Identifiers.ISBN10[0]
		}
	}

	return OLTitle, OLAuthor, OLPubDate, OLISBN, true
}