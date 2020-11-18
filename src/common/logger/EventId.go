package logger

/*
	EventIds are numeric representations of well-defined, intentional messages to the log analyst.
	We expect each EventId to be enriched with the FacilityId and various context-specific tags.
*/
type EventId uint32

const (
	// Core Logger operation EventId 0-999
	EventStd       EventId = 0 // Standard EventId for all normal log emissions.
	EventInit      EventId = 1 // Logging client is initialized.
	EventTagCreate EventId = 2 // Create tag event (descriptor).
	EventTagClose  EventId = 3 // Close tag event (descriptor).

	EventFlagForce EventId = 10 // Force flag

	// Reserved EventIds for tool and common use 1000+

)
