package logger

import "sync"

/*
	TagTracker keeps track of the tags assigned at any given time within the system.  This includes a tag dictionary
	of all tags known to the system as well as a `global` map of tags which will apply to all events so long as the
	tag is in this `global` table.
*/
const (
	tagPattern = `[a-zA-Z][a-zA-Z0-9]{0,63}`
	maxTagTrackerDictionarySize = 16384
)

type TagTracker struct {
	lock     sync.Mutex
	nextTag  TagId            // Next tagId to issue.
	global   map[TagId]bool   // This is the set of globally applied tags (they appear in all events until removed).
	tagNames map[string]TagId // Dictionary of tagNames and TagIds currently known to the system.
	tagIds   map[TagId]string // Dictionary of TagIds in the system for quick forward lookup
}
