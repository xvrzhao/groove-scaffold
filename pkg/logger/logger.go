package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var zlogger zerolog.Logger

func init() {
	cw := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime}
	cw.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("- %s", i)
	}
	zlogger = zerolog.New(cw).With().Timestamp().Logger()
}

func Info(format string, a ...any) {
	zlogger.Info().Msg(fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	zlogger.Warn().Msg(fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	zlogger.Error().Msg(fmt.Sprintf(format, a...))
}

func Fatal(format string, a ...any) {
	zlogger.Fatal().Msg(fmt.Sprintf(format, a...))
}
