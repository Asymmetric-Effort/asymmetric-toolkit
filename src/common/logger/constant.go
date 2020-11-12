package logger

import "time"

const (
	facilityRegExPattern = `^[a-zA-Z][a-zA-Z0-9]+$`
	// <timestamp>[<Facility>](<Level>): <message>"
	LogFormat                      = "[%s][%s](%s): %s"
	logWriteInterval time.Duration = 1    //Flush every x log entries.
	logBufferSz      int           = 1024 //Size of the log buffer channel.
	defaultLoggerFacility string = "Logger"
	loggerSetupReadyMessage = "Logger::Setup() complete."
)