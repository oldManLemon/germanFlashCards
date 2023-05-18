package sqlutils

import (
	"database/sql"
	"log"

	"github.com/oldManLemon/germanFlashCards/structs"
	"github.com/oldManLemon/germanFlashCards/zlogs"
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
	logger := zlogs.SetupLogger()
	logger.Debug().Str("germanWord", c.WordGerman).Msg("Insert_data initiated")
	//Prepare the statement
	statement, err := db.Prepare("INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES (?, ?, ?);")
	if err != nil {
		// log.Fatal(err)
		logger.Error().Err(err).Msg("Can not prepare statement")
		// fmt.Println(err)
	} else {
		defer statement.Close()
		_, err = statement.Exec(c.Article, c.WordGerman, c.WordEnglish)
		if err != nil {
			logger.Fatal().Err(err).Msg("Can not executre statement. No word has being added see err")
			// log.Fatal(err)
		}

	}
	//TODO RETURN TRUE HERE
	logger.Info().Str("germanWord", c.WordGerman).Msg("Data  has being inserted")

}

func Delete_data(c Card) bool {
	logger := zlogs.SetupLogger()
	logger.Debug().Str("germanWord", c.WordGerman).Msg("Initiate delete")
	//I am returning a straigt bool. I will capture logs later, I will not return an error here
	//so we are just return true for success and false for failure.

	// Prepare the statement
	statement, err := db.Prepare("DELETE FROM ger_dict WHERE ger_word = ?;")
	if err != nil {
		logger.Error().Err(err).Msg("Can not prepare statement")
		// log.Fatal(err)
		return false
	}
	defer statement.Close()

	_, err = statement.Query(c.WordGerman)
	if err != nil {
		logger.Error().Err(err).Msg("Can not preform query statement: ")
		// log.Fatal(err)
		return false
	}
	// Log the query
	// log.Println(statement)

	// add the following to delete all words with the given German word
	_, err = statement.Exec(c.WordGerman)
	if err != nil {
		logger.Error().Err(err).Msg("Can not execute statement: ")

		return false
	}
	// Convert the word to a string to make check_word happy and preform a sanity check
	word_convert := c.WordGerman
	sanity, err := Check_word(word_convert)
	if err != nil {
		logger.Error().Err(err).Msg("Check_word failed to call from Delete_data()")
		// fmt.Println("Error")
		return false
	}
	if sanity {
		logger.Error().Str("germanWord", c.WordGerman).Msg("Was unable to complete request to delete  from Database. Word was still discovered in the list")
		return false //Do I just go again

	}
	logger.Info().Str("germanWord", c.WordGerman).Msg(" successfully removed from the Database")
	return true

}

func Check_word(word string) (bool, error) {
	logger := zlogs.SetupLogger()
	logger.Debug().Str("germanWord", word).Msg("initiating check_word function")

	// Prepare the statement
	statement, err := db.Prepare("SELECT ger_word FROM ger_dict WHERE ger_word = ?;")
	if err != nil {
		logger.Error().Err(err).Msg("Can not prepare statement")
		return false, err
	}
	// Query for the word
	// https://go.dev/doc/tutorial/database-access#single_row
	var gerWord string //for the db
	err = statement.QueryRow(word).Scan(&gerWord)
	if err == sql.ErrNoRows {
		//TODO LOG IT
		logger.Error().Err(err).Msg("Word was not found")
		return false, nil //no word found
	} else if err != nil {
		logger.Error().Err(err).Msg("Query Error occured")
		return false, err //something else has gone wrong
	}
	logger.Info().Str("germanWord", word).Msg("Word was found in Database")
	return true, nil //word found

}
