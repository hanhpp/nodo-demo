package lrus

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var Log = &customLogrus{}
var LogFilePrefix = "demo"

func InitLogrus() {
	initLogrusOutput()
}

func GetFileName() string {
	return fmt.Sprintf("%s-%s.log", LogFilePrefix, time.Now().Format("01-02-2006"))
}

func initLogrusOutput() {
	logger := logrus.New()

	// Check if logs folder exists, create if not
	logsPath := filepath.Join(".", "data/logs")
	err := os.MkdirAll(logsPath, 0777)
	if err != nil {
		fmt.Println("logs path Err:", err)
	}

	// Currently only log out to one file
	fileName := GetFileName()
	filePath := filepath.Join("data/logs", fileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Infof("Failed to open file err : %v", err)
	}

	// let's use multi writer for output here, both std out and file
	output := io.MultiWriter(file, os.Stdout)
	logger.SetOutput(output)

	// Show line number and function name
	logger.SetReportCaller(false)

	Log.Instance = logger
	Log.FileName = fileName
}

// Logger collects logging information at several levels
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
	Close()
	Rotate()
	CallFileInfo() string
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
}

type customLogrus struct {
	FileName string
	Instance *logrus.Logger
	mu       sync.Mutex
}

func (l *customLogrus) WithField(k string, v interface{}) *customLogrus {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField(k, v)
	l.mu.Unlock()
	return l
}

func (l *customLogrus) WithFields(fields logrus.Fields) *customLogrus {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithFields(fields)
	l.mu.Unlock()
	return l
}

// Debug logs a message using DEBUG as log level.
func (l *customLogrus) Debug(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Debugln(v)
	l.mu.Unlock()
}

// Info logs a message using INFO as log level.
func (l *customLogrus) Info(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Info(v)
	l.mu.Unlock()
}

// Warning logs a message using WARNING as log level.
func (l *customLogrus) Warning(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Warningln(v)
	l.mu.Unlock()
}

// Error logs a message using ERROR as log level.
func (l *customLogrus) Error(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Errorln(v)
	l.mu.Unlock()
}

// Critical logs a message using CRITICAL as log level.
func (l *customLogrus) Critical(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).WithField("trace", "CRITICAL").Errorln(v)
	l.mu.Unlock()
}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func (l *customLogrus) Fatal(v ...interface{}) {
	l.mu.Lock()
	l.Rotate()
	l.Instance.WithField("file", l.CallFileInfo()).Fatalln(v)
	l.mu.Unlock()
}

func (l *customLogrus) Close() {
	l.Instance = nil
	l.FileName = ""
}

func (l *customLogrus) Rotate() {
	if GetFileName() != l.FileName {
		initLogrusOutput()
	}
}

func (l *customLogrus) CallFileInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	} else {
		const projectDirName = "satoshi-games"
		index := strings.Index(file, projectDirName)
		if index != -1 {
			return fmt.Sprintf(`%s:%d`, file[index:], line)
		}

		return fmt.Sprintf(`%s:%d`, file, line)
	}
}

type SkipLogger struct{}

func (l SkipLogger) Info(msg string, keysAndValues ...interface{}) {
	// fmt.Println(msg)
}

func (l SkipLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	Log.Error(err, msg)
}
