package logger

import (
	"log"
	"testing"

	. "gopkg.in/check.v1"
)

type MockWriter struct {
	lastString string
}

func (w *MockWriter) Write(p []byte) (int, error) {
	w.lastString = string(p)
	return 0, nil
}
func Test(t *testing.T) {
	TestingT(t)
}

type LogSuite struct {
	mockLogger *Logger
	mockWriter *MockWriter
}

var _ = Suite(&LogSuite{})

func (s *LogSuite) SetUpTest(c *C) {
	w := MockWriter{}
	l := log.New(&w, "", 0)
	s.mockLogger = &Logger{
		console:    nil,
		fileLogger: l,
		file:       nil,
		level:      INFO,
	}
	s.mockWriter = &w
}

func (s *LogSuite) TearDownTest(c *C) {

}

func (s *LogSuite) TestNewLoggerConsole(c *C) {
	l, err := NewLogger("", INFO)
	c.Assert(err, IsNil)
	c.Assert(l, NotNil)
}

func (s *LogSuite) TestTrace(c *C) {
	s.mockLogger.level = TRACE
	s.mockLogger.Trace("trace message")
	lastString := s.mockWriter.lastString
	c.Assert(lastString, Matches, ".*TRACE - trace message\n")
}

func (s *LogSuite) TestTraceNoPrint(c *C) {
	s.mockLogger.level = INFO
	s.mockLogger.Trace("trace message")
	lastString := s.mockWriter.lastString
	c.Assert(lastString, Equals, "")
}

func (s *LogSuite) TestDebug(c *C) {
	s.mockLogger.level = DEBUG
	s.mockLogger.Debug("debug message")
	lastString := s.mockWriter.lastString
	c.Assert(lastString, Matches, ".*DEBUG - debug message\n")
}

func (s *LogSuite) TestDebugLower(c *C) {
	s.mockLogger.level = TRACE
	s.mockLogger.Debug("debug message")
	lastString := s.mockWriter.lastString
	c.Assert(lastString, Matches, ".*DEBUG - debug message\n")
}

func (s *LogSuite) TestDebugNoPrint(c *C) {
	s.mockLogger.level = ERROR
	s.mockLogger.Debug("test123")
	lastString := s.mockWriter.lastString
	c.Assert(lastString, Equals, "")
}

func (s *LogSuite) TestInfo(c *C) {
	s.mockLogger.Info("test123")
	lastString := s.mockWriter.lastString
	c.Assert(lastString, Matches, ".*INFO - test123\n")
}

func (s *LogSuite) TestInfoNoPrint(c *C) {
	s.mockLogger.level = ERROR
	s.mockLogger.Info("test123")
	lastString := s.mockWriter.lastString
	c.Assert(lastString, Equals, "")
}
