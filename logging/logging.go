package logging

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
	logger = &Logger{
		c

	}
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

func (logger *Logger) Error(s string, err error) {
	logString := createLogString(Error, s, err)
	logger.logToConsole(logString)
}

func (logger *Logger) createLogString(level LogLevel, s string) string {
	logTime = time.Now()
	return fmt.Sprintf("%s | %s | %s", logTime, level, s)
}

func (logger *Logger) createLogString(level LogLevel, s string, err error) string {
	logTime = time.Now()
	return fmt.Sprintf("%s | %s | %s | %s", logTime, level, s, err.Error())
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
