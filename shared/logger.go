package shared

import (
	"fmt"
	"io"
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
	logrus      interface{}
	write       io.Writer
	debugMode   bool
}

func NewLogger(fileName ...interface{}) *logger {
	var list []string
	for _, v := range fileName {
		list = append(list, v.(string))
	}
	l := new(logger)
	l.debugMode = true
	// l.fileName = strings.Join(list, "-")

	if len(list) <= 0 {
		l.fileName = GetEnvOrFail("SYSTEM_NAME")
	}
	if len(l.fileName) <= 0 {
		l.fileName = "cron-go"
	}
	debugEnv := os.Getenv("DEGUB_MODE")
	if len(debugEnv) >= 1 {
		if debugEnv == "true" {
			l.debugMode = true
		} else if debugEnv == "false" {
			l.debugMode = false
		}

	}

	systemName := os.Getenv("SYSTEM_NAME")
	if len(systemName) <= 0 {
		systemName = "CRON-GO"
	}
	l.fields = logrus.Fields{"file_name": l.fileName, "system": systemName}

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
		if err != nil {
			fmt.Println("LOGGER: ERRO ON OPEN FILE", l.__dirname+"/../.logs/"+l.fileName+".log", err)
		}
		logrus.SetOutput(file)
		l.file = file
		// defer file.Close
		l.initialized = true
	}
}

func (l *logger) Info(text ...interface{}) {
	if l.initialized == false {
		l.init()
	}
	var ntext []interface{}
	ntext = append(ntext, "INFO: ")
	for _, t := range text {
		ntext = append(ntext, t)
	}
	logrus.WithFields(l.fields).Info(text...)
	if l.debugMode {
		cli_colors.PrintColor(cli_colors.Green, ntext...)
	}
}

func (l *logger) CloseFile() {
	l.file.Close()
	l.initialized = false
}

func (l *logger) Warning(text ...interface{}) {

	if l.initialized == false {
		l.init()
	}
	var ntext []interface{}
	ntext = append(ntext, "WARNING: ")
	for _, t := range text {
		ntext = append(ntext, t)
	}
	logrus.WithFields(l.fields).Warn(ntext...)
	if l.debugMode {
		cli_colors.PrintColor(cli_colors.Yellow, ntext...)
	}

}

func (l *logger) Write(data []byte) (int, error) {
	n, err := l.write.Write(data)
	if err != nil {
		return n, err
	}
	if n != len(data) {
		return n, io.ErrShortWrite
	}
	return len(data), nil
}

func (l *logger) Error(text ...interface{}) {

	if l.initialized == false {
		l.init()
	}
	var ntext []interface{}
	ntext = append(ntext, "ERROR: ")
	for _, t := range text {
		ntext = append(ntext, t)
	}
	logrus.WithFields(l.fields).Error(ntext...)
	if l.debugMode {
		cli_colors.PrintColor(cli_colors.Red, ntext...)
	}
}

func (l *logger) Fatal(text ...interface{}) {

	if l.initialized == false {
		l.init()
	}
	var ntext []interface{}
	ntext = append(ntext, "FATAL: ")
	for _, t := range text {
		ntext = append(ntext, t)
	}
	if l.debugMode {
		cli_colors.PrintColor(cli_colors.Red, ntext...)
	}
	logrus.WithFields(l.fields).Fatal(ntext...)
	os.Exit(1)
}
