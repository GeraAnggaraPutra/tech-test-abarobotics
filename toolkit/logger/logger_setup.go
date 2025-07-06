package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"abarobotics-test/toolkit/constant"
	"abarobotics-test/toolkit/util"

	"github.com/rs/zerolog"
)

type loggerOption struct {
	log   zerolog.Logger
	level int
}

var defaultLogger *loggerOption

func NewLogger() {
	var (
		level  = constant.DefaultLogLevel
		maxDir = constant.DefaultLogCallerMaxDirectory
	)

	level = util.ParseInt(level, os.Getenv("LOG_LEVEL"))
	maxDir = util.ParseInt(maxDir, os.Getenv("LOG_MAX_DIRECTORY"))

	cw := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("\n%s", i)
		},
		FormatCaller: func(i interface{}) string {
			dir, file := filepath.Split(fmt.Sprintf("%s", i))
			list := strings.Split(dir, "/")

			if len(list) < maxDir {
				return fmt.Sprintf("%s%s", dir, file)
			}

			return fmt.Sprintf("%s%s", strings.Join(list[len(list)-maxDir:], "/"), file)
		},
	}

	log := zerolog.New(cw).
		Level(parseLevelZerolog(level)).
		With().
		Timestamp().
		CallerWithSkipFrameCount(constant.DefaultLogCallerSkipFrame).
		Logger()

	defaultLogger = &loggerOption{
		log:   log,
		level: level,
	}
}
