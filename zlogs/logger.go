package zlogs

import (
	"log"
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
)

func SetupLogger() zerolog.Logger {
	//Setup File
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	buildInfo, _ := debug.ReadBuildInfo()
	//*Output to Console Only
	// logger := zerolog.New(zerolog.ConsoleWriter{
	// 	Out:        os.Stderr,
	// 	TimeFormat: time.RFC3339,
	// }).
	// 	Level(zerolog.TraceLevel).
	// 	With().
	// 	Timestamp().
	// 	Caller().
	// 	Int("pid", os.Getpid()).
	// 	Str("go_version", buildInfo.GoVersion).
	// 	Logger()

	//* Output to Logfile only
	// logger := zerolog.New(file).
	// 	Level(zerolog.TraceLevel).
	// 	With().
	// 	Timestamp().
	// 	Caller().
	// 	Int("pid", os.Getpid()).
	// 	Str("go_version", buildInfo.GoVersion).
	// 	Logger()

	//*Output to both

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	multiWriter := zerolog.MultiLevelWriter(consoleWriter, file)

	logger := zerolog.New(multiWriter).
		Level(zerolog.TraceLevel). //! Here Change the level!
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()

	return logger
}
