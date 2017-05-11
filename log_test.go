package logger

import (
	"log"
	"testing"

	. "gopkg.in/check.v1"
)

type mockLogWriter struct {
	lastString string
}

func (w *mockLogWriter) Write(p []byte) (int, error) {
	w.lastString = string(p)
	return 0, nil
}
func Test(t *testing.T) {
	TestingT(t)
}

type LogSuite struct {
	mockLogger        *Logger
	mockLogWriter     *mockLogWriter
	mockConsoleWriter *mockLogWriter
}

var _ = Suite(&LogSuite{})

func (s *LogSuite) SetUpTest(c *C) {
	w := mockLogWriter{}
	cw := mockLogWriter{}
	l := log.New(&w, "", 0)
	cl := log.New(&cw, "", 0)
	s.mockLogger = &Logger{
		console:    cl,
		fileLogger: l,
		file:       nil,
		level:      INFO,
	}
	s.mockLogWriter = &w
	s.mockConsoleWriter = &cw
}

func (s *LogSuite) TearDownTest(c *C) {

}

func (s *LogSuite) TestNewLoggerConsole(c *C) {
	l, err := NewLogger("", INFO)
	c.Assert(err, IsNil)
	c.Assert(l, NotNil)
}

func (s *LogSuite) TestConsole(c *C) {
	s.mockLogger.Console("console message")
	lastString := s.mockConsoleWriter.lastString
	c.Assert(lastString, Equals, "console message\n")
}

func (s *LogSuite) TestConsoleInfo(c *C) {
	s.mockLogger.ConsoleInfo("console info message")
	lastConsoleString := s.mockConsoleWriter.lastString
	lastLogString := s.mockLogWriter.lastString
	c.Assert(lastConsoleString, Equals, "console info message\n")
	c.Assert(lastLogString, Matches, ".*INFO - console info message\n")

}

func (s *LogSuite) TestTrace(c *C) {
	s.mockLogger.level = TRACE
	s.mockLogger.Trace("t1")
	lastString := s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*TRACE - t1\n")
}

func (s *LogSuite) TestTraceNoPrint(c *C) {
	var lastString string
	s.mockLogger.level = DEBUG
	s.mockLogger.Trace("t1")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = INFO
	s.mockLogger.Trace("t2")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = WARNING
	s.mockLogger.Trace("t3")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = ERROR
	s.mockLogger.Trace("t4")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = FATAL
	s.mockLogger.Trace("t5")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
}

func (s *LogSuite) TestDebug(c *C) {
	var lastString string
	s.mockLogger.level = DEBUG
	s.mockLogger.Debug("d1")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*DEBUG - d1\n")
	s.mockLogger.level = TRACE
	s.mockLogger.Debug("d2")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*DEBUG - d2\n")
}

func (s *LogSuite) TestDebugNoPrint(c *C) {
	var lastString string
	s.mockLogger.level = INFO
	s.mockLogger.Debug("d1")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = WARNING
	s.mockLogger.Debug("d2")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = ERROR
	s.mockLogger.Debug("d3")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = FATAL
	s.mockLogger.Debug("d4")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
}

func (s *LogSuite) TestInfo(c *C) {
	var lastString string
	s.mockLogger.level = TRACE
	s.mockLogger.Info("i1")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*INFO - i1\n")
	s.mockLogger.level = DEBUG
	s.mockLogger.Info("i2")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*INFO - i2\n")
	s.mockLogger.level = INFO
	s.mockLogger.Info("i3")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*INFO - i3\n")
}

func (s *LogSuite) TestInfoNoPrint(c *C) {
	var lastString string
	s.mockLogger.level = WARNING
	s.mockLogger.Info("i1")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = ERROR
	s.mockLogger.Info("i2")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = FATAL
	s.mockLogger.Info("i3")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
}

func (s *LogSuite) TestWarning(c *C) {
	var lastString string
	s.mockLogger.level = TRACE
	s.mockLogger.Warning("w1")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*WARNING - w1\n")
	s.mockLogger.level = DEBUG
	s.mockLogger.Warning("w2")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*WARNING - w2\n")
	s.mockLogger.level = INFO
	s.mockLogger.Warning("w3")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*WARNING - w3\n")
	s.mockLogger.level = WARNING
	s.mockLogger.Warning("w4")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*WARNING - w4\n")
}

func (s *LogSuite) TestWarningNoPrint(c *C) {
	var lastString string
	s.mockLogger.level = ERROR
	s.mockLogger.Warning("w1")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
	s.mockLogger.level = FATAL
	s.mockLogger.Warning("w2")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
}

func (s *LogSuite) TestError(c *C) {
	var lastString string
	s.mockLogger.level = TRACE
	s.mockLogger.Error("e1")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*ERROR - e1\n")
	s.mockLogger.level = DEBUG
	s.mockLogger.Error("e2")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*ERROR - e2\n")
	s.mockLogger.level = INFO
	s.mockLogger.Error("e3")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*ERROR - e3\n")
	s.mockLogger.level = WARNING
	s.mockLogger.Error("e4")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*ERROR - e4\n")
}

func (s *LogSuite) TestErrorNoPrint(c *C) {
	var lastString string
	s.mockLogger.level = FATAL
	s.mockLogger.Error("e1")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
}

func (s *LogSuite) TestClose(c *C) {
	err := s.mockLogger.Close()
	c.Assert(err, IsNil)
	c.Assert(s.mockLogger.file, IsNil)
}
