package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Card struct {
	article     string
	wordGerman  string
	wordEnglish string
}

func extractor(word string) (string, string) {
	/*
		Accepts only one argument to rerieve the word.
		Information will then be return, currently the Gender will be returned as either m,f,n and the second return will be the english translation of the word.

	*/

	//Capitalize First letter is needed
	title := strings.Title(word)

	url := fmt.Sprintf("https://de.wiktionary.org/wiki/%s", title)
	// Make HTTP GET request to the Wiktionary page

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Load HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the first <em> tag with a "title" attribute containing the word "Genus"
	em := doc.Find("em[title*=Genus]").First()
	if (em.Text()) == "" {
		fmt.Println("It was empty")
	}

	//Get English translation
	span := doc.Find("span[lang*=en]").First()
	eng := span.Find("a").Text()

	// <span lang="en"><a href="/wiki/table#table_(Englisch)" title="table">table</a></span>

	// Extract the grammatical gender text from the <em> tag
	gender := em.Text()

	// Print the grammatical gender
	// fmt.Println(gender)
	// fmt.Println(eng)
	return gender, eng
}
func NewCard(word string) Card {
	gender, eng := extractor(word)
	return Card{wordGerman: word, wordEnglish: eng, article: gender}
}

func main() {

	fmt.Println(NewCard("Tisch"))
	fmt.Println(NewCard("Tür"))
	fmt.Println(NewCard("Vogel"))
	fmt.Println(NewCard("Baum"))
	fmt.Println(NewCard("bier"))
	fmt.Println(NewCard("banane"))
	fmt.Println(NewCard("leben"))
	fmt.Println(NewCard("laufen"))
	fmt.Println(NewCard("Pferd"))
}
