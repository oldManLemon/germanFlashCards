package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/oldManLemon/germanFlashCards/functions"
)

func main() {
	//* Setup Logging
	// logger := zlogs.SetupLogger()
	// 	logger.Trace().Msg("trace message")
	// 	logger.Debug().Msg("debug message")
	// 	logger.Info().Msg("info message")
	// 	logger.Warn().Msg("warn message")
	// 	logger.Error().Msg("error message")
	// 	logger.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	// 	logger.WithLevel(zerolog.PanicLevel).Msg("panic message")

	// 	c := structs.Card{Article: "das", WordGerman: "Tisch", WordEnglish: "time"}
	// 	d := sqlutils.Delete_data(c)
	// 	fmt.Println(d)
	// logger.Info().Msg("Start Function Call!")
	// b, w1, w2 := extractor.Extractor("Tisch")
	// fmt.Println(b)
	// fmt.Println(w1)
	// fmt.Println(w2)
	b := functions.New_word("papier")
	fmt.Println(b)
}
