package functions

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3" //force the drive issue
	"github.com/oldManLemon/germanFlashCards/extractor"
	"github.com/oldManLemon/germanFlashCards/sqlutils"
	"github.com/oldManLemon/germanFlashCards/structs"
	"github.com/oldManLemon/germanFlashCards/zlogs"
)

func newCard(word string) (structs.Card, error) {
	logger := zlogs.SetupLogger()
	logger.Debug().Str("germanWord", word).Msg("initiating newCard function")
	found, gender, eng := extractor.Extractor(word)
	if !found {
		logger.Error().Str("germanWord", word).Msg("unable to to create a card, word is not found or does not contain article")
		return structs.Card{}, fmt.Errorf("could not create card, information extraction failed")
	}
	card := structs.Card{
		WordGerman:  word,
		WordEnglish: eng,
		Article:     gender,
	}
	logger.Info().Str("germanWord", card.WordGerman).Str("English Word", card.WordEnglish).Str("Article", card.Article).Msg("New card successfully created")
	return card, nil
}

func New_word(word string) string {
	logger := zlogs.SetupLogger()
	logger.Debug().Str("germanWord", word).Msg("initiating New_word function")

	check, err := sqlutils.Check_word(word)
	if !check { // faster way of having check == false
		logger.Debug().Bool("check value", check).Msg("Check returned false. Word will now attempt to be added")
		if err != nil {
			logger.Error().Err(err).Msg("Checking the word in the DB returned an error in the DB. ")
			msg := fmt.Sprintf("%v error", err.Error())
			return msg
		}
		//* Create a new word card, this will automatically call the extractor function
		new, err := newCard(word)
		//Handle error again
		if err != nil {
			logger.Error().Err(err).Msg("Extraction failed")
			msg := fmt.Sprintf("%v", err.Error())
			// fmt.Println(err)
			return msg
		} else {
			logger.Info().Msg("")
			fmt.Println(word)
			// sqlite_insert_data(test)
			sqlutils.Insert_data(new)
			logger.Info().Msg("Data insterted")
			msg := fmt.Sprintf("%v successfully added to collection\n", word)
			return msg
		}
	} else {
		logger.Debug().Str("germanWord", word).Msg("Already existed in collection")
		msg := fmt.Sprintf("%v already exists in collection\n", word)
		return msg
	}
	// return ""
}
