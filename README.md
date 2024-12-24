# goLogger

`goLogger` is a Go package that provides a flexible and efficient logging mechanism. It supports logging to files with optional buffering and offers integration with multiple timezones for precise timestamping.

## Features

- Logs messages at various levels: `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`.
- Buffered logging to improve performance.
- Supports rotation of log files with customizable backup count.
- Configurable timezone for timestamps.
- Graceful handling of fatal logs.

## Installation

To install the package, run:
```bash
go get github.com/sumit-behera-in/goLogger
```

## Usage

### Importing the Package
```go
import "github.com/sumit-behera-in/goLogger"
```

### Creating a Logger Instance
To create a logger instance, use `NewLogger`:
```go
logger, err := goLogger.NewLogger("MyApp", "./logs", 100, 5, "IST")
if err != nil {
    log.Fatalf("Failed to create logger: %v", err)
}
defer logger.Close()
```

Parameters:
- `appName`: Name of your application.
- `logFileDir`: Directory for log files. Leave empty to log to stdout.
- `bufferSize`: Number of logs to buffer before flushing to file.
- `logBackupCount`: Number of rotated log files to keep.
- `timeZone`: Timezone for timestamps (see [Supported Timezones](#supported-timezones)).

### Logging Messages
```go
logger.Debug("This is a debug message")
logger.Infof("Starting server on port %d", 8080)
logger.Warn("Low disk space")
logger.Error("Failed to connect to database")
logger.Fatal("Critical error. Shutting down.")
```

### Supported Timezones
The following timezones are supported for timestamping:

- **ACST**: Australia/Adelaide
- **AEST**: Australia/Sydney
- **AKST**: America/Anchorage
- **AST**: Asia/Riyadh
- **AWST**: Australia/Perth
- **BST**: Europe/London
- **CCT**: Asia/Shanghai
- **CDT**: America/Chicago
- **CET**: Europe/Paris
- **CST**: America/Chicago
- **EAT**: Africa/Nairobi
- **EDT**: America/New_York
- **EET**: Europe/Bucharest
- **EST**: America/New_York
- **GMT**: Europe/London
- **HKT**: Asia/Hong_Kong
- **HST**: Pacific/Honolulu
- **IST**: Asia/Calcutta
- **JST**: Asia/Tokyo
- **KST**: Asia/Seoul
- **MDT**: America/Denver
- **MSK**: Europe/Moscow
- **MST**: America/Denver
- **NZST**: Pacific/Auckland
- **PDT**: America/Los_Angeles
- **PST**: America/Los_Angeles
- **SAST**: Africa/Johannesburg
- **SGT**: Asia/Singapore
- **UTC**: UTC
- **WAT**: Africa/Lagos

If an invalid or unsupported timezone is provided, the logger defaults to the system's local timezone.

## Example
```go
package main

import (
    "github.com/sumit-behera-in/goLogger"
)

func main() {
    logger, err := goLogger.NewLogger("ExampleApp", "./logs", 100, 3, "UTC")
    if err != nil {
        panic(err)
    }
    defer logger.Close()

    logger.Info("Application started")
    logger.Debug("This is a debug log")
    logger.Warn("Warning: High memory usage")
    logger.Error("An error occurred")
}
```

## License
This project is licensed under the MIT License.
