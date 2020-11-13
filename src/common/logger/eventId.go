package logger
/*
	EventIds are numeric representations of well-defined, intentional messages to the log analyst.
	We expect each EventId to be enriched with the FacilityId and various context-specific tags.
 */
type EventId uint32

const (
	EventNoop      EventId = 0 //Unspecified event.
	EventInit      EventId = 1
	EventPanic     EventId = 2 //This indicates a general Panic augmented by tags.
	EventError     EventId = 3 //A general error
	EventTagCreate EventId = 4 //Create tag event
	EventTagClose  EventId = 5 //Close tag event
)
