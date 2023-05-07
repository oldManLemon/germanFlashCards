package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/oldManLemon/germanFlashCards/extractor"
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

	bool1, str1, str2 := extractor.Extractor("skdjfhsdf")

	fmt.Println(bool1)
	fmt.Println(str1)
	fmt.Println(str2)

}
