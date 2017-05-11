# go-log
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
To create a new logger
```go
log := NewLogger("/tmp/mylogfile.log", logger.INFO)
```
This will create an INFO level logger that will log to the path
`/tmp/mylogfile.log`.

To log to info:
```go
log.INFO("My log message")
```

To log to error:
```go
log.ERROR("My error message")
```
