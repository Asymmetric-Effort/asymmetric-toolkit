package logger

import "time"

const (
	logWriteInterval time.Duration = 1    //Flush every x log entries.
	logBufferSz      int           = 1024 //Size of the log buffer channel.
)