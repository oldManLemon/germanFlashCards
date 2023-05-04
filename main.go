package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	// "github.com/oldManLemon/germanFlashCards/sqlutils"
	"github.com/oldManLemon/germanFlashCards/extractor"
	"github.com/oldManLemon/germanFlashCards/structs"
)

func NewCard(word string) structs.Card {
	gender, eng := extractor.Extractor(word)
	card := structs.Card{
		WordGerman:  word,
		WordEnglish: eng,
		Article:     gender,
	}
	return card
}

func main() {

	test := NewCard("Schrank")
	fmt.Println(test)
	// sqlite_insert_data(test)
}
