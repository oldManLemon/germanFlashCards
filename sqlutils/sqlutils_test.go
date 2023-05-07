package sqlutils

import (
	"testing"

	_ "github.com/mattn/go-sqlite3" //force the drive issue
)

// THis annoyingly doesn't work :(
func TestNewWord(t *testing.T) {

	SetDbPathForTesting("../sqlite/ger_dict.db")
	// sqlite/ger_dict.db

	// Test for an existing word
	wordExists, err := Check_word("Schrank")

	if err != nil {
		t.Errorf("Error while checking word: %v", err)
	}
	if !wordExists {
		t.Errorf("Expected Schrank to exist, but it doesn't")
	}

	// Test for a non-existing word
	wordExists, err = Check_word("NonExistentWord")
	if err != nil {
		t.Errorf("Error while checking word: %v", err)
	}
	if wordExists {
		t.Errorf("Expected NonExistentWord to not exist, but it does")
	}

}
