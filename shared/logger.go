package shared

import (
	"fmt"
	"log"
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
}

func NewLogger(fileName string) *logger {
	l := new(logger)
	l.fileName = fileName
	return l
}

func (l *logger) init() {

	if l.initialized != true {
		if _, file, _, ok := runtime.Caller(0); ok {
			l.__dirname = path.Dir(file)
		} else {
			log.Fatal("error locating logger directory")
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
		} else {
			fmt.Println("ERRO ON FILE", err)
		}

		// defer file.Close

		l.fields = logrus.Fields{"file_name": l.fileName}
		l.initialized = true
	}
}

func (l *logger) Warning(text string) {

	if l.initialized == false {
		l.init()
	}
	logrus.WithFields(l.fields).Warn(text)
	cli_colors.PrintColor(cli_colors.Yellow, "WARNING: "+text)

}

func (l *logger) Error(text string) {

	if l.initialized == false {
		l.init()
	}
	logrus.WithFields(l.fields).Error(text)
	cli_colors.PrintColor(cli_colors.Red, "Error: "+text)
}

func (l *logger) Info(text string) {

	if l.initialized == false {
		l.init()
	}
	logrus.WithFields(l.fields).Info(text)
	cli_colors.PrintColor(cli_colors.Green, "Info: "+text)
}
