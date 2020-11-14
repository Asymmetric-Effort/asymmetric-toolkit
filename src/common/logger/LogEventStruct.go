package logger

import (
	"time"
)

type LogEventStruct struct {
	eventId EventId
	time    time.Time
	level   Level
	tags    []TagId
	message string
}