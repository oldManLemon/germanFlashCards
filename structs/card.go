package structs

type Card struct {
	Article     string
	WordGerman  string
	WordEnglish string
}

// To export we need to chnage the article and other field names to capitals. https://go.dev/ref/spec#Exported_identifiers
