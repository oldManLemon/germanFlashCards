package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/oldManLemon/germanFlashCards/functions"
	"github.com/rs/zerolog"
)

func main() {
	//* Setup Logging
	logger := functions.SetupLogger()
	logger.Trace().Msg("trace message")
	logger.Debug().Msg("debug message")
	logger.Info().Msg("info message")
	logger.Warn().Msg("warn message")
	logger.Error().Msg("error message")
	logger.WithLevel(zerolog.FatalLevel).Msg("fatal message")
	logger.WithLevel(zerolog.PanicLevel).Msg("panic message")
}
