package logger

type LogEventStruct struct {
	EventId EventId `json: "EventId"`
	Time    int64   `json: "Time"`
	Level   Level   `json: "Level"`
	Tags    []TagId `json: "Tags"`
	Message string  `json: "Message"`
}
