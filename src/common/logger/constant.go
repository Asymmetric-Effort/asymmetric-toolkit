package logger

import "time"

const (
	// Format:
	// <timestamp>[<Facility>](<level>): <message>.
	LogFormat                             = "[%s][%s](%s): %s"
	LogWriteInterval        time.Duration = 1    //Flush every x log entries.
	LogBufferSz             int           = 1024 //Size of the log buffer channel.
	DefaultLoggerFacility   string        = "Logger"
	LoggerSetupReadyMessage               = "Logger::Setup() complete."
)
