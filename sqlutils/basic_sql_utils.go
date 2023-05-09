package sqlutils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/oldManLemon/germanFlashCards/structs"
)

type Card = structs.Card

// Adding for test switches?
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

func Insert_data(c Card) {
	// // Open Sqlite Database file
	// db, err := sql.Open("sqlite3", "sqlite/ger_dict.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	// // Create a new table to store data
	// _, err = db.Exec("CREATE TABLE ger_dict (id INTERGER PRIMARY KEY, ger_article TEXT NOT NULL, ger_word TEXT NOT NULL UNIQUE, eng_word TEXT NOT NULL);")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//Prepare the statement
	statement, err := db.Prepare("INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES (?, ?, ?);")
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	} else {
		defer statement.Close()
		_, err = statement.Exec(c.Article, c.WordGerman, c.WordEnglish)
		if err != nil {
			log.Fatal(err)
		}

	}

}

func Delete_data(gerWord string) {
	// // Open Sqlite Database file
	// db, err := sql.Open("sqlite3", "sqlite/ger_dict.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// Prepare the statement
	statement, err := db.Prepare("DELETE FROM ger_dict WHERE ger_word = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	// Execute the statement
	_, err = statement.Exec(gerWord)
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
