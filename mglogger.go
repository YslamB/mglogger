package mglogger

import (
	"io"
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
	logFile, err := os.OpenFile(filePath+"/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

	if err != nil {
		panic(err)
	}

	if mode == "release" {
		l.SetOutput(logFile)

	} else {
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		l.SetOutput(multiWriter)
	}

	l.SetLevel(logrus.TraceLevel)
	return &Logger{logrus.NewEntry(l)}
}
