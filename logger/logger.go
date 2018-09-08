package logger

import (
	"log"
	"os"
)

// Logger ...
type Logger struct {
	logger *log.Logger
}

// New returns a configured instance of Logger.
func New(prefix string) *Logger {
	return &Logger{logger: log.New(os.Stderr, prefix+" ", log.LstdFlags)}
}
