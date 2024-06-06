package main

import (
	"errors"
	"fmt"
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

type LogHook struct {
}

// list of log levels that would trigger fire
func (hook *LogHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}

func (hook *LogHook) Fire(entry *logrus.Entry) error {
	if entry.Level == logrus.ErrorLevel {
		fmt.Println("ERROR: ", entry.Message)
		return errors.New("log error: " + entry.Message)
	}
	fmt.Println("WARNING: ", entry.Message)
	return nil
}

func TestLogHook(t *testing.T) {
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
	log.AddHook(&LogHook{})

	log.Info("Log outputted to log file")
	log.Warn("Unauthorized login detected")
	log.Error("Internal Server Error")
}
