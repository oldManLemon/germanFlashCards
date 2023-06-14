package functions

import (
	"fmt"

	"github.com/oldManLemon/germanFlashCards/sqlutils"
)

func Weight_random() []string {
	// I want to only return 10 words. If there are less than 10 words then return all the words.
	var test []string
	var words_and_scores map[string]float64

	words_and_scores, _ = sqlutils.Get_all_words()

	fmt.Println(words_and_scores)
	if len(words_and_scores) > 10 {
		// Get the words out of the words and score
		words := make([]string, 0, len(words_and_scores))
		for word := range words_and_scores { // for-each in x
			words = append(words, word)
		}
		return words
	} else {
		//Math!
	}

	return test

}
