package mglogger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func GetLogger(filePath string, fileName string, mode string) *Logger {

	l := logrus.New()
	l.SetReportCaller(true)

	l.SetFormatter(&logrus.JSONFormatter{})

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, 0777)
		if err != nil {
			panic(err)
		}
	}

	if mode == "release" {

		logFile, err := os.OpenFile(filePath+"/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}

		l.SetOutput(logFile)
	}

	l.SetLevel(logrus.TraceLevel)

	return &Logger{logrus.NewEntry(l)}
}
