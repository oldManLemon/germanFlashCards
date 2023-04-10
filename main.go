package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/oldManLemon/germanFlashCards/extractor"
)

type Card struct {
	article     string
	wordGerman  string
	wordEnglish string
}

func NewCard(word string) Card {
	gender, eng := extractor.Extractor(word)
	return Card{wordGerman: word, wordEnglish: eng, article: gender}
}

// DB Stuff

func sqlite_insert_data(c Card) {
	// Open Sqlite Database file
	db, err := sql.Open("sqlite3", "sqlite/ger_dict.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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
	}
	defer statement.Close()
	_, err = statement.Exec(c.article, c.wordGerman, c.wordEnglish)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	// test := NewCard("Katze")
	// fmt.Println(test)
	// sqlite_insert_data(test)

}
