package go_logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger struct holds the log format and the application name
type Logger struct {
	appName      string
	logger       *log.Logger
	timeLocation *time.Location // IST location
}

// NewLogger creates a new instance of Logger
func NewLogger(appName string, logFilePath string) (*Logger, error) {
	// Load IST location (Indian Standard Time)
	timeLocation, err := time.LoadLocation("Asia/Calcutta")
	if err != nil {
		return nil, fmt.Errorf("failed to load IST location: %v", err) // handle error if location is not found
	}

	var logOutput *os.File
	if logFilePath != "" {
		logOutput, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %v", err)
		}
	} else {
		logOutput = os.Stdout
	}

	return &Logger{
		appName:      appName,
		logger:       log.New(logOutput, "", 0), // Default to no timestamp
		timeLocation: timeLocation,
	}, nil
}

// logMessage prints a message with a log level and timestamp
func (l *Logger) logMessage(levelName string, msg string) {
	// Get current time in IST
	timestamp := time.Now().In(l.timeLocation).Format("2006-01-02 15:04:05") + " IST"
	logMsg := fmt.Sprintf("%s [%s] %s: %s", timestamp, levelName, l.appName, msg)
	l.logger.Println(logMsg)
}

// logMessagef prints a formatted message with a log level and timestamp
func (l *Logger) logMessagef(levelName string, format string, args ...interface{}) {
	// Get current time in IST
	timestamp := time.Now().In(l.timeLocation).Format("2006-01-02 15:04:05") + " IST"
	logMsg := fmt.Sprintf("%s [%s] %s: %s", timestamp, levelName, l.appName, fmt.Sprintf(format, args...))
	l.logger.Println(logMsg)
}

// Debug logs a message at DEBUG level
func (l *Logger) Debug(msg string) {
	l.logMessage("DEBUG", msg)
}

// Debugf logs a formatted message at DEBUG level
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logMessagef("DEBUG", format, args...)
}

// Info logs a message at INFO level
func (l *Logger) Info(msg string) {
	l.logMessage("INFO", msg)
}

// Infof logs a formatted message at INFO level
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logMessagef("INFO", format, args...)
}

// Warn logs a message at WARN level
func (l *Logger) Warn(msg string) {
	l.logMessage("WARN", msg)
}

// Warnf logs a formatted message at WARN level
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logMessagef("WARN", format, args...)
}

// Error logs a message at ERROR level
func (l *Logger) Error(msg string) {
	l.logMessage("ERROR", msg)
}

// Errorf logs a formatted message at ERROR level
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logMessagef("ERROR", format, args...)
}

// Fatal logs a message at FATAL level and terminates the program
func (l *Logger) Fatal(msg string) {
	l.logMessage("FATAL", msg)
	os.Exit(1)
}

// Fatalf logs a formatted message at FATAL level and terminates the program
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logMessagef("FATAL", format, args...)
	os.Exit(1)
}
