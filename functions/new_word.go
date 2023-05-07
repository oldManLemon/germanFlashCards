package functions

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3" //force the drive issue
	"github.com/oldManLemon/germanFlashCards/extractor"
	"github.com/oldManLemon/germanFlashCards/sqlutils"
	"github.com/oldManLemon/germanFlashCards/structs"
)

func newCard(word string) (structs.Card, error) {
	found, gender, eng := extractor.Extractor(word)
	if !found {
		return structs.Card{}, fmt.Errorf("could not create card, information extraction failed")
	}
	card := structs.Card{
		WordGerman:  word,
		WordEnglish: eng,
		Article:     gender,
	}
	return card, nil
}

func New_word(word string) string {

	check, err := sqlutils.Check_word(word)
	if !check { // faster way of having check == false
		if err != nil {
			msg := fmt.Sprintf("%v error", err.Error())
			return msg
		}
		new, err := newCard(word)
		//Handle error again
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(word)
			// sqlite_insert_data(test)
			sqlutils.Insert_data(new)
			msg := fmt.Sprintf("%v successfully added to collection\n", word)
			return msg
		}
	} else {
		msg := fmt.Sprintf("%v already exists in collection\n", word)
		return msg
	}
	return ""
}
