package logger

type Logger struct {
	Level       LogLevel
	Facility    Facility
	Destination Destination
	Writer      func(*string)
}
