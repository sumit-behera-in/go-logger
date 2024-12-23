package goLogger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// Logger struct for the logger
type Logger struct {
	appName      string
	logger       *log.Logger
	timeLocation *time.Location
	logBuffer    []string
	logFilePath  string
	bufferSize   int
	useBuffer    bool
	mu           sync.RWMutex // Protects logger state (non-buffer operations)
	bufferMu     sync.Mutex   // Protects logBuffer for buffered operations
	bufferCount  int
	flushCount   int
}

// NewLogger creates a new logger instance
func NewLogger(appName, logFilePath string) (*Logger, error) {
	timeLocation, err := time.LoadLocation("Asia/Calcutta")
	if err != nil {
		return nil, fmt.Errorf("failed to load IST location: %v", err)
	}

	var logOutput *os.File
	useBuffer := false
	if logFilePath != "" {
		// Check if the log file exists, and delete it if it does
		if _, err := os.Stat(logFilePath); err == nil {
			err := os.Remove(logFilePath)
			if err != nil {
				return nil, fmt.Errorf("failed to delete existing log file: %v", err)
			}
		}

		defer logOutput.Close()

		logOutput, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %v", err)
		}
		useBuffer = true
	} else {
		logOutput = os.Stdout
	}

	return &Logger{
		appName:      appName,
		logger:       log.New(logOutput, "", 0),
		timeLocation: timeLocation,
		logBuffer:    []string{},
		logFilePath:  logFilePath,
		bufferSize:   1000,
		bufferCount:  0,
		useBuffer:    useBuffer,
	}, nil
}

// logMessage writes a log message with a level
func (l *Logger) logMessage(levelName, msg string) {
	l.mu.RLock() // Allow concurrent reads
	defer l.mu.RUnlock()

	timestamp := time.Now().In(l.timeLocation).Format("2006-01-02 15:04:05") + " IST"
	logMsg := fmt.Sprintf("%s [%s] %s: %s", timestamp, levelName, l.appName, msg)

	if l.useBuffer && levelName != "FATAL" { // Don't buffer fatal messages
		l.addToBuffer(logMsg)
	}
	l.logger.Println(logMsg)
}

// addToBuffer safely appends a message to the buffer
func (l *Logger) addToBuffer(msg string) {
	l.bufferMu.Lock()
	defer l.bufferMu.Unlock()

	l.bufferCount++
	l.logBuffer = append(l.logBuffer, msg)
	if l.bufferCount >= l.bufferSize {
		l.bufferCount = 0
		go l.flushBuffer() // Flush in a separate goroutine to avoid blocking
	}
}

// flushBuffer safely writes buffered logs to the file
func (l *Logger) flushBuffer() {
	l.bufferMu.Lock()
	defer l.bufferMu.Unlock()

	if len(l.logBuffer) == 0 {
		return
	}

	// Generate the timestamped new file name
	renamedFilePath := l.logFilePath + fmt.Sprint(l.flushCount) + ".log"
	l.flushCount++

	// Close any open file before renaming
	// No need to open the file for renaming; just close the previous file handle.
	if l.logger != nil && l.logger.Writer() != nil {
		// Close the existing file handle
		if file, ok := l.logger.Writer().(*os.File); ok {
			file.Close()
		}
	}

	// Attempt to rename the log file
	err := os.Rename(l.logFilePath, renamedFilePath)
	if err != nil {
		log.Printf("Error renaming log file: %v", err)
		return
	}

	// Confirm that the file has been renamed
	println("Renamed log file to:", renamedFilePath)

	// Reopen the original log file for appending new logs
	file, err := os.OpenFile(l.logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
		return
	}
	defer file.Close()

	// Create a new logger with the updated file
	l.logger = log.New(file, "", 0)

	// Write buffered log messages to the new file
	for _, msg := range l.logBuffer {
		_, err := file.WriteString(msg + "\n")
		if err != nil {
			log.Printf("Error writing to log file: %v", err)
			return
		}
	}

	// Clear the log buffer after writing
	l.logBuffer = []string{}
}

// logMessagef writes a formatted log message with a level
func (l *Logger) logMessagef(levelName, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.logMessage(levelName, msg)
}

// fatalFlush writes a fatal log message directly to `fatal-log.log`
func (l *Logger) fatalFlush(msg string) {
	fatalLogFile := "fatal-log.log"
	file, err := os.OpenFile(fatalLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error opening fatal log file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(msg + "\n")
	if err != nil {
		fmt.Printf("Error writing to fatal log file: %v\n", err)
	}
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

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(msg string) {
	l.mu.Lock() // Lock to ensure thread-safe fatal logging
	defer l.mu.Unlock()

	timestamp := time.Now().In(l.timeLocation).Format("2006-01-02 15:04:05") + " IST"
	fatalMsg := fmt.Sprintf("%s [FATAL] %s: %s", timestamp, l.appName, msg)

	// Log fatal directly without buffering
	l.logger.Println(fatalMsg)

	// Write the fatal message to the dedicated log file
	l.fatalFlush(fatalMsg)

	os.Exit(1) // Terminate program
}

// Fatalf logs a formatted fatal message and exits
func (l *Logger) Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Fatal(msg)
}

// Close flushes any remaining buffer
func (l *Logger) Close() {
	if l.useBuffer {
		l.flushBuffer()
	}
}
