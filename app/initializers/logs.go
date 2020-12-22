package initializers

import (
	"strconv"

	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog"
)

const (
	LogLevelEnv     = "LOG_LEVEL"
	DefaultLogLevel = 0
)

func InitializeLogs() error {
	logLevelStr := envy.Get(LogLevelEnv, "0")

	logLevel, err := strconv.Atoi(logLevelStr)

	if err != nil {
		logLevel = DefaultLogLevel
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.Level(logLevel))

	return nil
}
