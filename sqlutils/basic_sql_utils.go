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
		println(err)
		log.Fatal(err)
	}
}

func Insert_data(c Card) {
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

func Delete_data(c Card) bool {
	//I am returning a straigt bool. I will capture logs later, I will not return an error here
	//so we are just return true for success and false for failure.

	// Prepare the statement
	statement, err := db.Prepare("DELETE FROM ger_dict WHERE ger_word = ?;")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer statement.Close()

	_, err = statement.Query(c.WordGerman)
	if err != nil {
		log.Fatal(err)
		return false
	}
	// Log the query
	// log.Println(statement)

	// add the following to delete all words with the given German word
	_, err = statement.Exec(c.WordGerman)
	if err != nil {
		fmt.Println("Error occured executing the staement: ", err.Error())
		log.Fatal(err)
		return false
	}
	// Convert the word to a string to make check_word happy and preform a sanity check
	word_convert := c.WordGerman
	sanity, err := Check_word(word_convert)
	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
		return false
	}
	if sanity {
		fmt.Println("Word in the list damnit ")
		return false

	}
	return true

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
		//TODO LOG IT
		// fmt.Println("Here no word found")
		return false, nil //no word found
	} else if err != nil {
		return false, err //something else has gone wrong
	}
	fmt.Println("found word yay!")
	return true, nil //word found

}
