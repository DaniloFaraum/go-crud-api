package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := os.Stdout                               //Defines the writer as the deafault OS output
	logger := log.New(writer, p, log.Ldate|log.Ltime) //Creates a logger with date and time in each log entry
}
