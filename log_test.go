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
	s.mockLogger.level = DEBUG
	s.mockLogger.Trace("t1")
	lastString := s.mockLogWriter.lastString
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
}

func (s *LogSuite) TestDebug(c *C) {
	s.mockLogger.level = DEBUG
	s.mockLogger.Debug("d1")
	lastString := s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*DEBUG - d1\n")
	s.mockLogger.level = TRACE
	s.mockLogger.Debug("d2")
	lastString = s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*DEBUG - d2\n")
}

func (s *LogSuite) TestDebugNoPrint(c *C) {
	s.mockLogger.level = ERROR
	s.mockLogger.Debug("test123")
	lastString := s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
}

func (s *LogSuite) TestInfo(c *C) {
	s.mockLogger.Info("test123")
	lastString := s.mockLogWriter.lastString
	c.Assert(lastString, Matches, ".*INFO - test123\n")
}

func (s *LogSuite) TestInfoNoPrint(c *C) {
	s.mockLogger.level = ERROR
	s.mockLogger.Info("test123")
	lastString := s.mockLogWriter.lastString
	c.Assert(lastString, Equals, "")
}
