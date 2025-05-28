package logger

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	once sync.Once
)

func InitLogger() {
	once.Do(func() {
		output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
		output.FormatLevel = func(i interface{}) string {
			if ll, ok := i.(string); ok {
				switch ll {
				case "debug":
					return "\033[36m[DBG]\033[0m"
				case "info":
					return "\033[32m[INF]\033[0m"
				case "warn":
					return "\033[33m[WRN]\033[0m"
				case "error":
					return "\033[31m[ERR]\033[0m"
				default:
					return "[" + ll + "]"
				}
			}
			return "[???]"
		}
		log.Logger = zerolog.New(output).With().Timestamp().Logger()
	})
}

// Info logs an informational message
func Info(message string) {
	log.Info().Msg(message)
}

// Error logs an error message with an error object
func Error(err error, message string) {
	log.Error().Err(err).Msg(message)
}

// Debug logs a debug message
func Debug(message string) {
	log.Debug().Msg(message)
}

// Warn logs a warning message
func Warn(message string) {
	log.Warn().Msg(message)
}

// Fatal logs a fatal message and exits the program
func Fatal(message string) {
	log.Fatal().Msg(message)
}

// Fatalf logs a fatal message and exits the program with a formatted message
func Fatalf(format string, v ...interface{}) {
	log.Fatal().Msgf(format, v...)
}
