package logger

import "sync"

/*
	TagTracker keeps track of the tags assigned at any given time within the system.  This includes a tag dictionary
	of all tags known to the system as well as a `global` map of tags which will apply to all events so long as the
	tag is in this `global` table.
*/


type TagTracker struct {
	lock     sync.Mutex
	nextTag  TagId             // Next tagId to issue.
	global   TagTable          // This is the set of globally applied tags (they appear in all events until removed).
	tagNames TagNameDictionary // Dictionary of tagNames and TagIds currently known to the system.
	tagIds   TagDictionary     // Dictionary of TagIds in the system for quick forward lookup
}



