package sqlutils

import (
	"database/sql"
	"fmt"
	"log"
)

// Adding for testing as we can't test from main.go and the paths are now wrong. Downside of sqlite.
const defaultDbPath = "sqlite/ger_dict.db"

var dbPath = defaultDbPath
var db *sql.DB

func SetDbPathForTesting(path string) {
	dbPath = path
}

func init() {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
}

func Check_word(word string) (bool, error) {

	// Prepare the statement
	statement, err := db.Prepare("SELECT ger_word FROM ger_dict WHERE ger_word = ?;")
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return false, err
	}
	// Query for the word
	// https://go.dev/doc/tutorial/database-access#single_row
	var gerWord string //for the db
	err = statement.QueryRow(word).Scan(&gerWord)
	if err == sql.ErrNoRows {
		return false, nil //no word found
	} else if err != nil {
		return false, err //something else has gone wrong
	}
	return true, nil //word found

}
