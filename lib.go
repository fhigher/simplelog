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

type LoggerWrap struct {
	out   io.WriteCloser
	inner *log.Logger
	level loglevel
	path  string
}

func newSimpleLogger(out io.WriteCloser) *LoggerWrap {
	return &LoggerWrap{
		out:   out,
		inner: log.New(out, "", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func StdLogger() *LoggerWrap {
	lw := newSimpleLogger(os.Stdout)
	lw.path = "stdout"

	return lw
}

func FileLogger(path string) (*LoggerWrap, error) {
	fobj, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if nil != err {
		return nil, err
	}

	lw := newSimpleLogger(fobj)
	lw.path = path
	return lw, nil
}

func (lw *LoggerWrap) GetPath() string {
	return lw.path
}

func (lw *LoggerWrap) SetLevel(level loglevel) {
	lw.level = level
}

func (lw *LoggerWrap) GetLevel() string {
	return logNameMap[lw.level]
}

func (lw *LoggerWrap) Close() error {
	return lw.out.Close()
}

func (lw *LoggerWrap) Trace(v ...interface{}) {
	if TRACE < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[TRACE])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *LoggerWrap) Tracef(format string, v ...interface{}) {
	if TRACE < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[TRACE])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}

func (lw *LoggerWrap) Debug(v ...interface{}) {
	if DEBUG < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[DEBUG])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *LoggerWrap) Debugf(format string, v ...interface{}) {
	if DEBUG < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[DEBUG])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}

func (lw *LoggerWrap) Info(v ...interface{}) {
	if INFO < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[INFO])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *LoggerWrap) Infof(format string, v ...interface{}) {
	if INFO < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[INFO])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}

func (lw *LoggerWrap) Warn(v ...interface{}) {
	if WARN < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[WARN])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *LoggerWrap) Warnf(format string, v ...interface{}) {
	if WARN < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[WARN])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}

func (lw *LoggerWrap) Error(v ...interface{}) {
	if ERROR < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[ERROR])
	lw.inner.Output(2, fmt.Sprint(v...))
}

func (lw *LoggerWrap) Errorf(format string, v ...interface{}) {
	if ERROR < lw.level {
		return
	}
	lw.inner.SetPrefix(logNameMap[ERROR])
	lw.inner.Output(2, fmt.Sprintf(format, v...))
}
