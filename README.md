# Go Logger

`goLogger` is a simple, flexible logging package for Go applications. It allows you to log messages at various log levels (DEBUG, INFO, WARN, ERROR, FATAL) with timestamps in **Indian Standard Time (IST)**. The log entries include the application name, log level, and message.

## Features

- **Custom Log Levels**: Supports multiple log levels such as DEBUG, INFO, WARN, ERROR, FATAL.
- **Timestamp in IST**: Logs are timestamped in **Indian Standard Time (IST)** with "IST" appended at the end.
- **Formatted Logging**: Supports both plain messages and formatted log messages.
- **Customizable Application Name**: The logger allows you to set a custom application name that will appear in each log entry.
- **Fatal Logging**: Logs a FATAL message and terminates the program execution.

## Installation

To use `goLogger` in your Go project, simply include it as part of your project structure. If you're using Go modules, you can use `go get` to fetch it.

1. Create a Go module or navigate to your existing Go project.
2. Run the following command to install the package:
   ```bash
   go get github.com/sumit-behera-in/goLogger
   ```

## Usage

### Initializing the Logger

To use the logger in your application, you need to create a new instance of the `Logger` with your application name. 

```go
package main

import (
	"log"
	"github.com/sumit-behera-in/goLogger"
)

func main() {
	// Initialize the logger with your application name
	logger, err := goLogger.NewLogger("MyApp")
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	// Log messages at various levels
	logger.Debug("This is a debug message")
	logger.Infof("This is an info message with a number: %d", 42)
	logger.Warnf("This is a warning message: %s", "be careful!")
	logger.Errorf("An error occurred: %s", "file not found")
	
	// Fatal log that terminates the program
	logger.Fatalf("This is a fatal error, exiting...")
}
```

### Log Levels

- **Debug**: Used for low-level debugging information.
- **Info**: Used for general informational messages.
- **Warn**: Used for warnings about potentially harmful situations.
- **Error**: Used for error messages indicating failures.
- **Fatal**: Logs critical errors and terminates the program.

Each of these log levels can be logged using the following methods:

- `Debug(msg string)`
- `Debugf(format string, args ...interface{})`
- `Info(msg string)`
- `Infof(format string, args ...interface{})`
- `Warn(msg string)`
- `Warnf(format string, args ...interface{})`
- `Error(msg string)`
- `Errorf(format string, args ...interface{})`
- `Fatal(msg string)`
- `Fatalf(format string, args ...interface{})`

### Example Output

Here's an example of the output format for each log level:

```
2024-12-22 15:34:45 IST [DEBUG] MyApp: This is a debug message
2024-12-22 15:34:45 IST [INFO] MyApp: This is an info message with a number: 42
2024-12-22 15:34:45 IST [WARN] MyApp: This is a warning message: be careful!
2024-12-22 15:34:45 IST [ERROR] MyApp: An error occurred: file not found
2024-12-22 15:34:45 IST [FATAL] MyApp: This is a fatal error, exiting...
```

### Error Handling

If there is an issue loading the IST location (e.g., incorrect time zone), the logger will return an error, which you should handle appropriately. Example:

```go
logger, err := goLogger.NewLogger("MyApp")
if err != nil {
	log.Fatalf("Error initializing logger: %v", err)
}
```

## Methods Overview

### `NewLogger(appName string) (*Logger, error)`

Creates a new instance of `Logger` with the provided application name and the default IST time zone.

- **Parameters**:
  - `appName`: The name of the application that will be included in each log entry.
- **Returns**:
  - A pointer to a `Logger` instance.
  - An error if the IST time zone could not be loaded.

### `Debug(msg string)`

Logs a message with a **DEBUG** level.

### `Debugf(format string, args ...interface{})`

Logs a formatted message with a **DEBUG** level.

### `Info(msg string)`

Logs a message with an **INFO** level.

### `Infof(format string, args ...interface{})`

Logs a formatted message with an **INFO** level.

### `Warn(msg string)`

Logs a message with a **WARN** level.

### `Warnf(format string, args ...interface{})`

Logs a formatted message with a **WARN** level.

### `Error(msg string)`

Logs a message with an **ERROR** level.

### `Errorf(format string, args ...interface{})`

Logs a formatted message with an **ERROR** level.

### `Fatal(msg string)`

Logs a message with a **FATAL** level and terminates the program.

### `Fatalf(format string, args ...interface{})`

Logs a formatted message with a **FATAL** level and terminates the program.