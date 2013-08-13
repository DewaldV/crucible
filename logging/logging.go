package logging

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type LoggerConfig struct {
	Level    LogLevel
	FileName string
}

func NewLoggerConfig() (c *LoggerConfig) {
	c = new(LoggerConfig)
	c.FileName = "log"
	c.Level = Info
	return
}

func NewLogger(c *LoggerConfig) (logger *Logger) {
	return
}

type LogType string
type LogLevel string

const (
	Error   LogLevel = "ERROR"
	Warning          = "WARN"
	Info             = "INFO"
	Debug            = "DEBUG"
)

type Logger struct {
	loggerConfig *LoggerConfig
	logger       *log.Logger
}

func (logger *Logger) Debug(s string) {
	logString := createLogString(Debug, s)
	logger.logToConsole(logString)
}

func (logger *Logger) Info(s string) {
	logString := createLogString(Info, s)
	logger.logToConsole(logString)
}

func (logger *Logger) Warning(s string) {
	logString := createLogString(Warning, s)
	logger.logToConsole(logString)
}

func (logger *Logger) Error(s string) {
	logString := createLogString(Error, s)
	logger.logToConsole(logString)
}

func createLogString(level LogLevel, s string) string {
	logTime := time.Now()
	return fmt.Sprintf("%s | %s | %s", logTime, level, s)
}

func (logger *Logger) logToConsole(logString string) {
	fmt.Println(logString)
	return
}

func (logger *Logger) logToFile(logString string) {

}

func LogRequest(request *http.Request) (s string) {
	s = fmt.Sprint("Method: %s > Request: %s\n", request.Method, request.URL.String())
	return
}
