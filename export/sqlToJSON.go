package export

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"projects/LibraryDatabase2.0/database"
)

// Reads SQL rows and reformulates the said rows into the JSON format.
// The JSON data is exported as a file.
func SqlToJson() bool {
	rows, err := database.DBCon.Query("SELECT id, author, title, publish_date, isbn FROM library ORDER BY author ASC")
	if err != nil {
		fmt.Println("There was an error: ", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("There was an error: ", err)
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		fmt.Println("There was an error: ", err)
	}

	writeError := ioutil.WriteFile("./books.json", jsonData, 0644)
	if writeError == nil {
		return true
	} else {
		fmt.Println("There was an error writing the file: ", writeError)
		return false
	}
}