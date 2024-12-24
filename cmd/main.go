package main

import (
	"github.com/sumit-behera-in/goLogger"
)

func main() {
	// Create a new logger
	logger, err := goLogger.NewLogger("test", "log", 80000, 40, "")
	if err != nil {
		panic(err.Error())
	}

	logger.Debugf("This is a formatted debug message with args: %s", "arg1")
	logger.Infof("This is a formatted info message with args: %s", "arg1")
	logger.Warnf("This is a formatted warn message with args: %s", "arg1")
	logger.Errorf("This is a formatted error message with args: %s", "arg1")

	logger.Debug("This is a debug message")
	logger.Info("This is a info message")
	logger.Warn("This is a warn message")
	logger.Error("This is a error message")

	defer logger.Close()

}
