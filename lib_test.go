package simplelog

import "testing"

func TestFileLogger(t *testing.T) {
	path := "test.log"
	lw, err := FileLogger(path)
	if nil != err {
		t.Fatal(err)
	}

	lw.SetLevel(WARN)

	lw.Trace("trace log")
	lw.Tracef("%s", "test trace log")

	lw.Debug("debug log")
	lw.Debugf("%s", "test debug log")

	lw.Info("info log")
	lw.Infof("%s", "test info log")

	lw.Warn("warn log")
	lw.Warnf("%s", "test warn log")

	lw.Error("error log")
	lw.Errorf("%s", "test error log")

	lw.Close()
}

func TestStdLogger(t *testing.T) {
	lw, err := StdLogger()
	if nil != err {
		t.Fatal(err)
	}

	lw.SetLevel(WARN)

	lw.Trace("trace log")
	lw.Tracef("%s", "test trace log")

	lw.Debug("debug log")
	lw.Debugf("%s", "test debug log")

	lw.Info("info log")
	lw.Infof("%s", "test info log")

	lw.Warn("warn log")
	lw.Warnf("%s", "test warn log")

	lw.Error("error log")
	lw.Errorf("%s", "test error log")

	lw.Close()
}
