package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/oldManLemon/germanFlashCards/sqlutils"
	"github.com/oldManLemon/germanFlashCards/structs"
	"github.com/oldManLemon/germanFlashCards/zlogs"
)

func main() {
	//* Setup Logging
	logger := zlogs.SetupLogger()
	// 	logger.Trace().Msg("trace message")
	// 	logger.Debug().Msg("debug message")
	// 	logger.Info().Msg("info message")
	// 	logger.Warn().Msg("warn message")
	// 	logger.Error().Msg("error message")
	// 	logger.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	// 	logger.WithLevel(zerolog.PanicLevel).Msg("panic message")

	c := structs.Card{Article: "das", WordGerman: "Tempo", WordEnglish: "time"}
	d := sqlutils.Delete_data()
}
