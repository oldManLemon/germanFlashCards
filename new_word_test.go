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

	//* Test case 2: Word doesn't exist
	word2 := "sjkdfhsdkjfhsdk"
	result2 := functions.New_word(word2)
	expectedOutput2 := ("could not create card, information extraction failed")
	if result2 != expectedOutput2 {
		t.Errorf("Test case 2 failed: Expected %s, BUT GOT %s", expectedOutput2, result2)
	}

	// expectedOutput := "Tisch already exists in collection\n"
	// if consoleOutput != expectedOutput {
	// 	t.Errorf("Test case 1 failed: Expected %s, but got %s", expectedOutput, consoleOutput)
	// }
}
