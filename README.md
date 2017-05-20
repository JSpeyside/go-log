# go-log
[![Build Status](https://travis-ci.org/jlgrady1/go-log.svg?branch=master)](https://travis-ci.org/jlgrady1/go-log) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jlgrady1/go-log)  [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/jlgrady1/go-log/master/LICENSE)
[![goreportcard](https://goreportcard.com/badge/github.com/jlgrady1/go-log](https://goreportcard.com/badge/github.com/jlgrady1/go-log)

Golang Logging Library

This Golang library provides an easy interface for logging to console and files.
The library provides support for multiple log levels.

## Supported Log levels
The following log levels are supported when logging.
* TRACE
* DEBUG
* INFO
* WARNING
* ERROR
* FATAL

Logging to fatal will terminate the program.

## Usage
To create a new logger specify a log file and a log level. Specifying an empty string `""` for the log file will return a console only logger.
```go
log := logger.NewLogger("/tmp/mylogfile.log", logger.INFO)
```
This will create an INFO level logger that will log to the path
`/tmp/mylogfile.log`.

To log to info:
```go
log.Info("My log message")
```

To log to error:
```go
log.Error("My error message")
```

You can even use formatting
```go
myString := "World"
myInt := 47
log.Debug("Hello, %s: %d", myString, myInt)
```
Yields `Hello, World: 47`
