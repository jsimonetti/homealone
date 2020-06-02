package log

import (
	"os"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

// Fields are a representation of formatted log fields
type Fields map[string]interface{}

// Logger is the interface for a logger
type Logger interface {
	logrus.StdLogger
	With(fields Fields) Logger
	WithError(err error) Logger
}

type logger struct {
	*logrus.Entry
}

// NewLogger returns a Logger based on logrus
func NewLogger() Logger {
	log := &struct {
		*logrus.Logger
	}{
		Logger: logrus.New(),
	}

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.0000"
	customFormatter.DisableColors = true
	customFormatter.FullTimestamp = true

	log.Formatter = customFormatter

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.Out = os.Stdout

	// Only log the debug severity or above.
	log.Level = logrus.DebugLevel

	// Disable concurrency mutex as we use Stdout
	log.SetNoLock()
	return &logger{Entry: log.WithFields(nil)}
}

// With will add the fields to the formatted log entry
func (l *logger) With(fields Fields) Logger {
	return &logger{Entry: l.WithFields(logrus.Fields(fields))}
}

// WithError will add the error to the fields
func (l *logger) WithError(err error) Logger {
	_, f, line, ok := runtime.Caller(1)
	file := strings.SplitAfter(f, "/homealone/")
	if ok {
		return l.With(Fields{"error": err, "file": file[1], "line": line})
	}
	return l.With(Fields{"error": err})

}
