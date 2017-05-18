package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

// LogLevel contains the range of logging levels to show
type LogLevel int

// Defines standard log levels
const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARNING //4
	ERROR
	FATAL
)

// Log is an interface describing the required functions for a logger.
type Log interface {
	Close() error
	Console(string, ...interface{})
	ConsoleInfo(string, ...interface{})
	Trace(string, ...interface{})
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warning(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
}

// Logger is a log handler for various levels of logging.
type Logger struct {
	console    *log.Logger
	fileLogger *log.Logger
	file       *os.File
	level      LogLevel
}

// Console logs a message directly to the console if quiet is not set.
func (logger *Logger) Console(message string, a ...interface{}) {
	msg := fmt.Sprintf(message, a...)
	logger.console.Println(msg)
}

// ConsoleInfo is a helper function that logs to both the console and Info log.
func (logger *Logger) ConsoleInfo(message string, a ...interface{}) {
	logger.Console(message, a...)
	logger.Info(message, a...)
}

// Trace writes a TRACE level log message to the log file.
func (logger *Logger) Trace(message string, a ...interface{}) {
	if logger.fileLogger == nil || logger.level > TRACE {
		return
	}
	msg := fmt.Sprintf(message, a...)
	logger.fileLogger.Println(logLine("TRACE", msg))
}

// Debug writes a DEBUG level log message to the log file.
func (logger *Logger) Debug(message string, a ...interface{}) {
	if logger.fileLogger == nil || logger.level > DEBUG {
		return
	}
	msg := fmt.Sprintf(message, a...)
	logger.fileLogger.Println(logLine("DEBUG", msg))
}

// Info writes an INFO level log message to the log file.
func (logger *Logger) Info(message string, a ...interface{}) {
	if logger.fileLogger == nil || logger.level > INFO {
		return
	}
	msg := fmt.Sprintf(message, a...)
	logger.fileLogger.Println(logLine("INFO", msg))
}

// Warning writes a WARNING level log message to the log file.
func (logger *Logger) Warning(message string, a ...interface{}) {
	if logger.fileLogger == nil || logger.level > WARNING {
		return
	}
	msg := fmt.Sprintf(message, a...)
	logger.fileLogger.Println(logLine("WARNING", msg))
}

// Error writes an ERROR level log message to the log file.
func (logger *Logger) Error(message string, a ...interface{}) {
	if logger.fileLogger == nil || logger.level > ERROR {
		return
	}
	msg := fmt.Sprintf(message, a...)
	logger.fileLogger.Println(logLine("ERROR", msg))
}

// Fatal writes a FATAL level log message to the log file.
func (logger *Logger) Fatal(message string, a ...interface{}) {
	if logger.fileLogger == nil || logger.level > FATAL {
		return
	}
	msg := fmt.Sprintf(message, a...)
	logger.fileLogger.Println(logLine("FATAL", msg))
	logger.Close()
	log.Fatalln(message)
}

// Close cleans up any file handles and necessary shutdown functions.
func (logger *Logger) Close() error {
	if logger.file == nil {
		return nil
	}
	return logger.file.Close()
}

func (logger *Logger) basicConfig() error {
	logger.console = log.New(os.Stdout, "", 0)
	logger.fileLogger = log.New(logger.file, "", 0)
	return nil
}

func logLine(level string, message string) string {
	//2017-05-09 18:26:00,211 - amqp - connection:755 - DEBUG -
	date := time.Now().Format("2006-01-02 15:04:05")
	_, filename, line, _ := runtime.Caller(2)
	filename = path.Base(filename)
	logLine := fmt.Sprintf("%s - %s:%d - %s - %s",
		date,
		filename,
		line,
		level,
		message,
	)
	return logLine
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			return false
		}
	}
	return true
}

func pathWritable(path string) bool {
	systemOS := runtime.GOOS
	if systemOS == "darwin" || systemOS == "linux" {
		return unix.Access(path, unix.W_OK) == nil
	}
	return true
}

// NewLogger returns a pointer to a new Logger containing basic configuration.
func NewLogger(filename string, level LogLevel) (Log, error) {
	// return a console logger if no filename is specified.
	if filename == "" {
		return &Logger{
			console:    log.New(os.Stdout, "", 0),
			fileLogger: nil,
			file:       nil,
			level:      level,
		}, nil
	}
	// Check the file path
	fileDir := filepath.Dir(filename)
	if !pathExists(fileDir) {
		return nil, fmt.Errorf("Log path does not exist %s", fileDir)
	}
	if !pathWritable(fileDir) {
		return nil, fmt.Errorf("Log path %s is not writable", fileDir)
	}

	// Create the file if it does not exist, otherwise open for appending
	var f *os.File
	var err error
	if pathExists(filename) {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return nil, errors.Wrapf(err, "Error opening file %s", filename)
		}
	} else {
		f, err = os.Create(filename)
		if err != nil {
			return nil, errors.Wrapf(err, "Error creating file %s", filename)
		}
	}
	logger := &Logger{
		console:    nil,
		fileLogger: nil,
		file:       f,
		level:      level,
	}

	logger.basicConfig()
	return logger, nil
}

// MockLogger is a dummy logger that performs no-ops.
type MockLogger struct{}

// NewMockLogger returns a new dummy logger.
func NewMockLogger() Log {
	return &MockLogger{}
}

// Close performs no-op.
func (logger *MockLogger) Close() error { return nil }

// Console performs no-op.
func (logger *MockLogger) Console(string, ...interface{}) {}

// ConsoleInfo performs no-op.
func (logger *MockLogger) ConsoleInfo(string, ...interface{}) {}

// Trace performs no-op.
func (logger *MockLogger) Trace(string, ...interface{}) {}

// Debug performs no-op.
func (logger *MockLogger) Debug(string, ...interface{}) {}

// Info performs no-op.
func (logger *MockLogger) Info(string, ...interface{}) {}

// Warning performs no-op.
func (logger *MockLogger) Warning(string, ...interface{}) {}

// Error performs no-op.
func (logger *MockLogger) Error(string, ...interface{}) {}

// Fatal performs no-op.
func (logger *MockLogger) Fatal(string, ...interface{}) {}
