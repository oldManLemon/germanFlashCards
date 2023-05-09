package main

import (
	"testing"

	_ "github.com/mattn/go-sqlite3" //force the drive issue
	"github.com/oldManLemon/germanFlashCards/functions"
)

func TestNew_word(t *testing.T) {
	// var consoleOutput string

	//* Test case 1: Word already exists
	word := "Tisch"
	result := functions.New_word(word)
	expectedOutput := (word + " already exists in collection\n")
	if result != expectedOutput {
		t.Errorf("Test case 1 failed: Expected %s, but got %s", expectedOutput, result)
	}

	//* Test case 2: Word doesn't exist and cannot be found
	word2 := "sjkdfhsdkjfhsdk"
	result2 := functions.New_word(word2)
	expectedOutput2 := ("could not create card, information extraction failed")
	if result2 != expectedOutput2 {
		t.Errorf("Test case 2 failed: Expected %s, BUT GOT %s", expectedOutput2, result2)
	}

	// //* Test case 3: Word doesn't exist, is found and successfully is added to collection
	//TODO create delete command.
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("failed to create mock database connection: %v", err)
	// }
	// defer db.Close()

	// // Set up the mock expectations
	// mock.ExpectPrepare("INSERT INTO ger_dict").ExpectExec().
	// 	WithArgs("die", "Polizei4", "Police4").
	// 	WillReturnResult(sqlmock.NewResult(1, 1))

	// // Call the function being tested
	// c := structs.Card{Article: "die", WordGerman: "Polizei4", WordEnglish: "Police4"}
	// sqlutils.Insert_data(c)

	// // Verify that the expected queries were executed
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %v", err)
	// }

	// expectedOutput := "Tisch already exists in collection\n"
	// if consoleOutput != expectedOutput {
	// 	t.Errorf("Test case 1 failed: Expected %s, but got %s", expectedOutput, consoleOutput)
	// }
}
