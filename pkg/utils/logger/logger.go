package logger

import (
	"github.com/sirupsen/logrus"
)

type LogMessageType string

const (
	JSON LogMessageType = "json"
	TEXT LogMessageType = "text"
)

var log *logrus.Logger = logrus.New()

func Configure(outputType LogMessageType, level logrus.Level) {
	switch outputType {
	case JSON:
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05", // Customize the timestamp format
		})
	case TEXT:
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,                  // Include the date and time in the timestamp
			TimestampFormat: "2006-01-02 15:04:05", // Customize the timestamp format
		})
	}
	log.SetReportCaller(true)
	log.SetLevel(level)
}

func Log() *logrus.Logger {
	return log
}
