package shared

import (
	"fmt"
	"os"
	"path"
	"runtime"

	cli_colors "github.com/sergiodii/cron-go/shared/cli/colors"
	"github.com/sirupsen/logrus"
)

type logger struct {
	fileName    string
	initialized bool
	fields      logrus.Fields
	__dirname   string
	file        *os.File
}

func NewLogger(fileName string) *logger {
	l := new(logger)
	l.fileName = fileName
	if len(fileName) <= 0 {
		l.fileName = GetEnvOrFail("SYSTEM_NAME")
	}
	return l
}

func (l *logger) init() {

	if l.initialized != true {
		if _, file, _, ok := runtime.Caller(0); ok {
			l.__dirname = path.Dir(file)
		} else {
			fmt.Println("error locating logger directory")
			return
		}
		logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "@timestamp",
				logrus.FieldKeyMsg:  "message",
			},
		})
		logrus.SetLevel(logrus.TraceLevel)

		file, err := os.OpenFile(l.__dirname+"/../.logs/"+l.fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			logrus.SetOutput(file)
			fmt.Println("ERRO ON FILE", err)
		}
		l.file = file
		// defer file.Close

		systemName := os.Getenv("SYSTEM_NAME")
		if len(systemName) <= 0 {
			systemName = "CRON-GO"
		}

		l.fields = logrus.Fields{"file_name": l.fileName, "system": systemName}
		l.initialized = true
	}
}

func (l *logger) Info(text ...interface{}) {

	if l.initialized == false {
		l.init()
	}
	var ntext []interface{}
	ntext = append(ntext, "INFO: ")
	for t := range text {
		ntext = append(ntext, t)
	}
	logrus.WithFields(l.fields).Info(ntext...)
	cli_colors.PrintColor(cli_colors.Green, ntext...)
}

func (l *logger) CloseFile() {
	l.file.Close()
}

func (l *logger) Warning(text ...interface{}) {

	if l.initialized == false {
		l.init()
	}
	var ntext []interface{}
	ntext = append(ntext, "WARNING: ")
	for t := range text {
		ntext = append(ntext, t)
	}
	logrus.WithFields(l.fields).Warn(ntext...)
	cli_colors.PrintColor(cli_colors.Yellow, ntext...)

}

func (l *logger) Error(text ...interface{}) {

	if l.initialized == false {
		l.init()
	}
	var ntext []interface{}
	ntext = append(ntext, "ERROR: ")
	for t := range text {
		ntext = append(ntext, t)
	}
	logrus.WithFields(l.fields).Error(ntext...)
	cli_colors.PrintColor(cli_colors.Red, ntext...)
}

func (l *logger) Fatal(text ...interface{}) {

	if l.initialized == false {
		l.init()
	}
	var ntext []interface{}
	ntext = append(ntext, "FATAL: ")
	for t := range text {
		ntext = append(ntext, t)
	}
	cli_colors.PrintColor(cli_colors.Red, ntext...)
	logrus.WithFields(l.fields).Fatal(ntext...)
	os.Exit(1)
}
