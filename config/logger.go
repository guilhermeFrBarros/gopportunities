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

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout) // cria um writer com o caminho padrão para escrita( o console)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
	}
}

// create no-formatter logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Err(v ...interface{}) {
	l.err.Println(v...)
}

// create Formatter logs Enabled Logs
func (l *Logger) DebugF(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) InfoF(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) WarningF(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
func (l *Logger) ErrF(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
