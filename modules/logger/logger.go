package loggo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/moatorres/go/modules/utils"
)

/*
Loggo is a JSON-based logger.
It uses native Go modules to parse messages and native or custom errors into JSON format.
Loggo also implements 'Fatal' that will print logs in JSON before calling os.Exit(1).

Usage:

	var logger = loggo.New(loggo.LoggerOptions{
		Service: "my-service",
	})

logger.Info("Hello 👋")
// → {"service":"my-service","level":"INFO","message":"Hello 👋","timestamp":"2023-10-12T10:15:18-03:00"}

	func (m *MyError) Error() string {
		return "Boom 💥"
	}

logger.Error("Unexpected Error: %s", err)
// → {"service":"my-service","level":"ERROR","message":"Unexpected Error: Boom 💥","timestamp":"2023-10-12T10:20:47-03:00"}

logger.Fatal("Uhoh 💀")
// → {"service":"my-service","level":"FATAL","message":"Uhoh 💀","timestamp":"2023-10-12T10:15:18-03:00"}
// exit status 1
*/
type Logger struct {
	Options LoggerOptions
}

type LoggerOptions struct {
	Service string
	Level   LogLevel
}

func New(options LoggerOptions) *Logger {
	return &Logger{
		Options: options,
	}
}

type LogMessage struct {
	Service   string `json:"service"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type LogLevel int

const (
	LOG LogLevel = iota
	INFO
	WARN
	ERROR
	DEBUG
	FATAL
)

func (l LogLevel) String() string {
	return [...]string{"LOG", "INFO", "WARN", "ERROR", "DEBUG", "FATAL"}[l]
}

type LogFunc func(string, ...interface{})

func (l *Logger) logWithLevel(level LogLevel, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)

	// create our log object
	log := LogMessage{
		Service:   l.Options.Service,
		Level:     level.String(),
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	// serialize to json
	json, err := json.Marshal(log)
	if err != nil {
		fmt.Printf("Loggo Error: %s\n", err.Error())
		return
	}

	// get level color
	color := getColor(level)
	// colorize
	output := utils.Colorize(string(json), []string{log.Service, log.Level, log.Message, log.Timestamp}, color)
	// print
	fmt.Println(string(output))
}

func (l *Logger) Log(format string, args ...interface{}) {
	l.logWithLevel(LOG, format, args...)
}
func (l *Logger) Info(format string, args ...interface{}) {
	l.logWithLevel(INFO, format, args...)
}
func (l *Logger) Debug(format string, args ...interface{}) {
	l.logWithLevel(DEBUG, format, args...)
}
func (l *Logger) Warn(format string, args ...interface{}) {
	l.logWithLevel(WARN, format, args...)
}
func (l *Logger) Error(format string, args ...interface{}) {
	l.logWithLevel(ERROR, format, args...)
}

// internal
func (l *Logger) fatal(format string, args ...interface{}) {
	l.logWithLevel(FATAL, format, args...)
}

func getColor(level LogLevel) utils.AnsiColor {
	switch level {
	case LOG:
		return utils.GREEN
	case INFO:
		return utils.BLUE
	case WARN:
		return utils.YELLOW
	case ERROR:
		return utils.LIGHT_RED
	case DEBUG:
		return utils.PURPLE
	case FATAL:
		return utils.RED
	default:
		return utils.RESET
	}
}

func (l *Logger) Fatal(args ...interface{}) {
	var msg string

	for _, arg := range args {
		switch v := arg.(type) {
		case error:
			msg = v.Error()
		case string:
			msg = v
		case LoggerOptions:
			l.Options = v
		}
	}

	l.fatal("%s", msg)

	os.Exit(1)
}
