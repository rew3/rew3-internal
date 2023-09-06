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

func Info(message ...interface{}) {
	log.Infoln(message...)
}

func Debug(message ...interface{}) {
	log.Debugln(message...)
}

func Warn(message ...interface{}) {
	log.Warn(message...)
}

func Println(message ...interface{}) {
	log.Println(message...)
}

func Error(message ...interface{}) {
	log.Error(message...)
}

func Fatal(message ...interface{}) {
	log.Fatal(message...)
}

func Panic(message ...interface{}) {
	log.Panic(message...)
}
