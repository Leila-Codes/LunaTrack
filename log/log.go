package log

import "github.com/sirupsen/logrus"

var logger *logrus.Logger
var newEntry *logrus.Entry

func init() {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	newEntry = logger.WithField("Application", "Luna-Track")
}

func GetLogger() *logrus.Entry {
	return newEntry
}
