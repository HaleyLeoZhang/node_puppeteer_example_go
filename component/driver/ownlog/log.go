package ownlog

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func Init(c *Config) {
	if STDOUT_YES == c.Stdout {
		if c.Dir == "" {
			log.Out = os.Stdout
		} else {
			file, err := os.OpenFile(c.Dir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				log.Out = file
			} else {
				log.Info("Failed to log to file, using default stderr")
			}
		}
	}
}

func Infof(format string, args ...interface{}) {
	log.Logf(logrus.InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Logf(logrus.WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Logf(logrus.ErrorLevel, format, args...)
}

func Info(args ...interface{}) {
	log.Log(logrus.InfoLevel, args...)
}

func Warn(args ...interface{}) {
	log.Log(logrus.WarnLevel, args...)
}

func Error(args ...interface{}) {
	log.Log(logrus.ErrorLevel, args...)
}
