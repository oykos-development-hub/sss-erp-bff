package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func Initialize() error {
	// Set the logger format to JSON for structured logging
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// Set logger level (this can be adjusted as needed: InfoLevel, WarnLevel, etc.)
	Logger.SetLevel(logrus.InfoLevel)

	// Open the log file for writing
	logFile, err := os.OpenFile("./log/sss-erp-bff.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Set the logger output to the log file
	Logger.SetOutput(logFile)

	return nil
}
