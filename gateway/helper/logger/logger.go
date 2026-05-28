package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"time"

	"github.com/sirupsen/logrus"
)

var Log *Logger
var mu sync.Mutex

type Fields logrus.Fields

type Logger struct {
	log *logrus.Entry
	logrusBase *logrus.Logger

}

func (l *Logger) BaseLog(fields map[string]interface{}, data interface{}) (log *logrus.Entry) {
	file, function, line := GetCaller()
	log = l.log.WithFields(logrus.Fields{
		"file":     file,
		"line":     line,
		"function": function,
	}).WithFields(fields)
	if data != nil {
		payload, err := json.Marshal(data)
		if err != nil {
			log = log.WithField("data_error", err.Error())
			return
		}
		log = log.WithField("data", string(payload))
	}
	return
}

func (l *Logger) Info(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Info(message)
}
func (l *Logger) Warn(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Warn(message)
}
func (l *Logger) Error(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Error(message)
}
func (l *Logger) Fatal(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Fatal(message)
}
func (l *Logger) Panic(fields map[string]interface{}, data interface{}, message string) {
	l.BaseLog(fields, data).Panic(message)
}

func (l *Logger) SetFileOutput(file io.Writer) {
	l.log.Logger.Out = file
}

func NewLogger(name string) {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log = &Logger{
		log: l.WithField("case", name),
		logrusBase: l,
	}

	go Log.dailyLogRotation(name)
}

func (l *Logger) dailyLogRotation(appName string) {
	for {
		l.rotateLogFile(appName)

		// Tunggu sampai pukul 00:00:01 hari berikutnya
		now := time.Now()
		next := now.Add(24 * time.Hour).Truncate(24 * time.Hour).Add(time.Second)
		time.Sleep(time.Until(next))
	}
}

func (l *Logger) rotateLogFile(appName string) {
	mu.Lock()
	defer mu.Unlock()

	dateStr := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s-%s.log", appName, dateStr)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to create/open log file: %v\n", err)
		return
	}

	
	l.logrusBase.SetOutput(file)
}
