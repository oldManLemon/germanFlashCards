package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getGender(word string) string {

	url := fmt.Sprintf("https://de.wiktionary.org/wiki/%s", word)
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

	// Extract the grammatical gender text from the <em> tag
	gender := em.Text()

	// Print the grammatical gender
	fmt.Println(gender)
	return gender

}

func main() {
	getGender("Tisch")
	getGender("Tür")
	getGender("Vogel")
	getGender("Baum")
	getGender("Apfel")
	getGender("banane")

}
