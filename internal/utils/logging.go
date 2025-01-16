package utils

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

// SetupLogger configures the logging level and format based on the provided input.
// It supports various log levels and two formats: plain text and JSON.
func SetupLogger(logLevel string, logFormat string) {
	setLogLevel(logLevel)
	setLogFormat(logFormat)

	log.WithFields(log.Fields{
		"log_level": log.GetLevel(),
		"formatter": logFormat,
	}).Debug("Logger configuration set")
}

// setLogLevel configures the logging level based on the provided input.
func setLogLevel(logLevel string) {
	switch strings.ToUpper(logLevel) {
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	default:
		log.Panic("No valid log level was set")
	}
}

// setLogFormat configures the logging format based on the provided input.
func setLogFormat(logFormat string) {
	switch strings.ToUpper(logFormat) {
	case "PLAIN":
		log.SetFormatter(&log.TextFormatter{
			ForceColors:            true,
			ForceQuote:             true,
			DisableLevelTruncation: true,
			PadLevelText:           true,
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
		})
	case "JSON":
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		log.Panic("No valid formatter was set")
	}
}
