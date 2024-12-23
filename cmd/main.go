package main

import (
	"sync"

	"github.com/sumit-behera-in/goLogger"
)

func main() {
	// Create a new logger
	logger, err := goLogger.NewLogger("test", "log.log")
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 800; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			logger.Debug("This is a debug message")
			logger.Debugf("This is a formatted debug message with args: %s", "arg1")
			logger.Info("This is an info message")
			logger.Infof("This is a formatted info message with args: %s", "arg1")
			logger.Warn("This is a warn message")
			logger.Warnf("This is a formatted warn message with args: %s", "arg1")
			logger.Error("This is an error message")
			logger.Errorf("This is a formatted error message with args: %s", "arg1")
		}()
	}

	wg.Wait()
	defer logger.Close()

}
