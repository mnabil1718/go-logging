package main

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutputLog(t *testing.T) {
	log := logrus.New()

	// create logs dir first
	err := os.MkdirAll("./logs/", 0755)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("./logs/application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(file)

	log.Info("Log outputted to log file")
	log.Warn("Log outputted to log file")
	log.Error("Log outputted to log file")
}

func TestLoggerWithField(t *testing.T) {
	log := logrus.New()

	// create logs dir first
	err := os.MkdirAll("./logs/", 0755)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("./logs/application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	log.SetOutput(file)

	log.WithField("loggedin_user", "mnabil1718").Info("Log outputted to log file")
	log.WithField("current_user", "mnabil1718").WithField("role", "admin").Info("Log outputted to log file")
	log.WithFields(logrus.Fields{
		"username":   "mnabil1718",
		"role":       "admin",
		"authorized": false,
	}).Warn("User unauthorized login")
}
