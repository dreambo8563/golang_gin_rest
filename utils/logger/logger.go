package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func init() {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}
	// log.WithField("project", "golang_gin_rest")
	logger = log.WithFields(logrus.Fields{
		"project": "golang_gin_rest",
	})
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

// Log return the custermized log instannce
func Log() *logrus.Entry {
	return logger
}
