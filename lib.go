package simplelog

import (
	"fmt"
	"io"
	"log"
	"os"
)

type loglevel int8

const (
	TRACE loglevel = iota
	DEBUG
	INFO
	WARN
	ERROR
)

var logNameMap = map[loglevel]string{
	TRACE: "TRACE ",
	DEBUG: "DEBUG ",
	INFO:  "INFO ",
	WARN:  "WARN ",
	ERROR: "ERROR ",
}

type loggerWrap struct {
	out   io.WriteCloser
	inner *log.Logger
	level loglevel
	path  string
}

func newSimpleLogger(out io.WriteCloser) *loggerWrap {
	return &loggerWrap{
		out:   out,
		inner: log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func StdLogger() (*loggerWrap, error) {
	lw := newSimpleLogger(os.Stdout)
	lw.path = "stdout"

	return lw, nil
}

func FileLogger(path string) (*loggerWrap, error) {
	fobj, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if nil != err {
		return nil, err
	}

	lw := newSimpleLogger(fobj)
	lw.path = path
	return lw, nil
}

func (lw *loggerWrap) GetPath() string {
	return lw.path
}

func (lw *loggerWrap) SetLevel(level loglevel) {
	lw.level = level
}

func (lw *loggerWrap) GetLevel() string {
	return logNameMap[lw.level]
}

func (lw *loggerWrap) Close() error {
	return lw.out.Close()
}

func (lw *loggerWrap) Trace(v ...interface{}) {
	if TRACE < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[TRACE])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *loggerWrap) Tracef(format string, v ...interface{}) {
	if TRACE < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[TRACE])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}

func (lw *loggerWrap) Debug(v ...interface{}) {
	if DEBUG < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[DEBUG])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *loggerWrap) Debugf(format string, v ...interface{}) {
	if DEBUG < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[DEBUG])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}

func (lw *loggerWrap) Info(v ...interface{}) {
	if INFO < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[INFO])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *loggerWrap) Infof(format string, v ...interface{}) {
	if INFO < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[INFO])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}

func (lw *loggerWrap) Warn(v ...interface{}) {
	if WARN < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[WARN])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *loggerWrap) Warnf(format string, v ...interface{}) {
	if WARN < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[WARN])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}

func (lw *loggerWrap) Error(v ...interface{}) {
	if ERROR < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[ERROR])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *loggerWrap) Errorf(format string, v ...interface{}) {
	if ERROR < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[ERROR])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}
