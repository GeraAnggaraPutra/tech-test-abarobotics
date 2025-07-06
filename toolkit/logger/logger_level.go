package logger

import (
	"github.com/rs/zerolog"
)

const (
	debugLevel = iota
	infoLevel
	warnLevel
	errorLevel
	fatalLevel
)

func parseLevelZerolog(level int) zerolog.Level {
	switch level {
	case debugLevel:
		return zerolog.DebugLevel
	case infoLevel:
		return zerolog.InfoLevel
	case warnLevel:
		return zerolog.WarnLevel
	case errorLevel:
		return zerolog.ErrorLevel
	case fatalLevel:
		return zerolog.FatalLevel
	}

	return zerolog.TraceLevel
}
