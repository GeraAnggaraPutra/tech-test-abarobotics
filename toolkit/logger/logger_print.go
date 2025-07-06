package logger

import "fmt"

// PrintJSON will print struct data as json string.
func PrintJSON(data ...interface{}) {
	for i := range data {
		fmt.Println(ParseJSON(data[i]))
	}
}

// PrettyPrint will print struct data as json indent string.
func PrettyPrint(data ...interface{}) {
	for i := range data {
		fmt.Println(ParsePrettyJSON(data[i]))
	}
}

// Print debug message.
func PrintDebug(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)
	defaultLogger.log.Debug().Msg(msg)
}

// Print info message.
func PrintInfo(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)
	defaultLogger.log.Info().Msg(msg)
}

// Print warn message.
func PrintWarn(msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)
	defaultLogger.log.Warn().Msg(msg)
}

// Print error message.
func PrintError(err error, msg string, fields ...interface{}) {
	msg = generateMessage(msg, fields)
	defaultLogger.log.Error().Err(err).Msg(msg)
}

// Print error message.
func PrintNewError(err, newError error, fields ...interface{}) error {
	msg := generateMessage(newError.Error(), fields)
	defaultLogger.log.Error().Err(err).Msg(msg)

	return newError
}

// Print fatal message.
func PrintFatal(err, newError error, fields ...interface{}) {
	msg := generateMessage(newError.Error(), fields)
	defaultLogger.log.Fatal().Err(err).Msg(msg)
}
