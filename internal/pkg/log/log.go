// Package log provides a default logger for the library.
package log

import "log"

// Logger is a struct that holds the logger.
type Logger struct{}

// NewLogger returns a new Logger.
func NewLogger() *Logger {
	return &Logger{}
}

// Debug logs the given message at the debug level.
func (l *Logger) Debug(msg string) {
	log.Println(msg)
}
