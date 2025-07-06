package logger

import (
	"context"

	"github.com/rs/zerolog"
)

type Logger struct {
	log   zerolog.Logger
	level int
}

func WithContext(ctx context.Context) *Logger {
	l := Logger{
		log:   defaultLogger.log,
		level: defaultLogger.level,
	}

	l.log.WithContext(ctx)

	return &l
}

// Debug logs debug messages.
func (l *Logger) Debug(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= debugLevel {
		l.log.Debug().Msg(msg)
	}
}

// Info logs info messages.
func (l *Logger) Info(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= infoLevel {
		l.log.Info().Msg(msg)
	}
}

// Warn logs warning messages.
func (l *Logger) Warn(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= warnLevel {
		l.log.Warn().Msg(msg)
	}
}

// Error logs error messages.
func (l *Logger) Error(err error, msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= errorLevel {
		l.log.Error().Err(err).Msg(msg)
	}
}

// Error logs error messages and return new error.
func (l *Logger) NewError(err, newErr error, fields ...interface{}) error {
	msg := generateMessage(newErr.Error(), fields)

	if l.level <= errorLevel {
		l.log.Error().Err(err).Msg(msg)
	}

	return newErr
}

// Fatal logs fatal messages.
func (l *Logger) Fatal(err error, msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)

	if l.level <= fatalLevel {
		l.log.Fatal().Err(err).Msg(msg)
	}
}
