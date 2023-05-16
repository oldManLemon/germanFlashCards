package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/oldManLemon/germanFlashCards/sqlutils"
	"github.com/oldManLemon/germanFlashCards/structs"
)

// func NewCard(word string) structs.Card {
// 	gender, eng := extractor.Extractor(word)
// 	card := structs.Card{
// 		WordGerman:  word,
// 		WordEnglish: eng,
// 		Article:     gender,
// 	}
// 	return card
// }

func main() {

	c := structs.Card{Article: "die", WordGerman: "Tempo", WordEnglish: "Police5"}
	d := sqlutils.Delete_data(c)
	fmt.Println(d)

}
