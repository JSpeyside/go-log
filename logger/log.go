package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger is a log handler for various levels of logging.
type Logger struct {
	console *log.Logger
	trace   *log.Logger
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	quiet   bool
	file    *os.File
}

// Console logs a message directly to the console if quiet is not set.
func (log Logger) Console(message string) {
	if log.quiet == true {
		return
	}
	log.console.Println(message)
}

// Info writes an INFO level log message to the log file.
func (log Logger) Info(message string) error {
	log.info.Println(message)
	return nil
}

func (log Logger) Close(message string) error {
	return log.file.Close()
}

func (logger Logger) Log(message string) error {
	// now := time.Now()
	// fmt.Printf("%s %s\n", now.Format(time.RFC3339), message)
	fmt.Println(message)
	return nil
}

func (logger Logger) LogLines(messages []string) error {
	for _, message := range messages {
		logger.Log(message)
	}
	return nil
}

func NewLogger(quiet bool, file string) (*Logger, error) {
	// file, err := os.Create(logFilePath)
	if err != nil {
		return nil, err
	}

	// consoleHandler = os.Stdout

	logger := &Logger{
		console: log.New(os.Stdout, "", 0),
		trace:   nil,
		debug:   nil,
		info:    nil,
		warning: nil,
		err:     nil,

		file:  file,
		quiet: quiet,
	}
	logger.info = log.New(logger.file, "INFO: ", 0)
	return logger, nil
}
