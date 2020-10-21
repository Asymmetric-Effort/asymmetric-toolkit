package logger

import "time"

const (
	// <timestamp>[<Facility>](<level>): <message>"
	LogFormat                      = "[%s][%s](%s): %s"
	logWriteInterval time.Duration = 1    //Flush every x log entries.
	logBufferSz      int           = 1024 //Size of the log buffer channel.
	defaultLoggerFacility string = "Logger"
	loggerSetupReadyMessage = "Logger::Setup() complete."
)
