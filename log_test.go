package logger

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type LogSuite struct {
}

var _ = Suite(&LogSuite{})
