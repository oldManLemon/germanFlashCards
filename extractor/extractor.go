package extractor

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/oldManLemon/germanFlashCards/zlogs"
)

func Extractor(word string) (bool, string, string) {
	/*
		Accepts only one argument to rerieve the word.
		Information will then be return, currently the Gender will be returned as either m,f,n and the second return will be the english translation of the word.

	*/
	logger := zlogs.SetupLogger()
	logger.Debug().Str("Word", word).Msg("Initiate word extrator")

	//Capitalize First letter is needed
	title := strings.Title(word)

	url := fmt.Sprintf("https://de.wiktionary.org/wiki/%s", title)
	logger.Debug().Str("URL", url).Msg("URL searched")
	// fmt.Println(url)
	// Make HTTP GET request to the Wiktionary page
	client := http.Client{
		Timeout: 20 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {

		//Die here straight away if word not found
		// fmt.Println(err)
		logger.Error().Err(err).Msg("Waiting for response error")
		return false, "", ""
	}
	defer resp.Body.Close()

	// Load HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Error().Err(err).Msg("Something has failed in the resp.Body parsing")
		return false, "", ""
	}

	// Find the first <em> tag with a "title" attribute containing the word "Genus"
	em := doc.Find("em[title*=Genus]").First()
	if (em.Text()) == "" {
		logger.Error().Err(err).Msg("Article was Empty")
		// fmt.Println("It was empty")
		return false, "", ""
	} else {
		//Get English translation
		span := doc.Find("span[lang*=en]").First()
		eng := span.Find("a").Text()
		// <span lang="en"><a href="/wiki/table#table_(Englisch)" title="table">table</a></span>
		// Extract the grammatical gender text from the <em> tag
		logger.Info().Str("word", word).Msg("Word information successfully extracted")
		gender := em.Text()
		return true, gender, eng

	}

}
