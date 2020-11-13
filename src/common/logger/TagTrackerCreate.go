package logger
/*
	TagTracker::Create() will create a new tag using the given tag name
	and return an associated numeric TagId.
 */
import (
	"regexp"
	"strings"
)

func (o *TagTracker) Create(rawTagName *string) TagId {
	re := regexp.MustCompile(tagPattern)
	if rawTagName == nil {
		panic("logger/TagTracker::Create() encountered nil tagName")
	}
	if o.tagNames == nil {
		o.tagNames = make(map[string]TagId,1)
	}
	if o.tagIds == nil {
		o.tagIds = make(map[TagId]string,1)
	}
	if len(o.tagIds) >= maxTagTrackerDictionarySize {
		panic("Too many tags are currently open.  Create failed.")
	}

	id:=o.nextTag
	tagName:= strings.TrimSpace(*rawTagName)

	if re.MatchString(tagName) {
		o.lock.Lock()
		defer o.lock.Unlock()
		o.tagIds[id] = tagName
		o.tagNames[tagName] = id
		o.nextTag++
	} else {
		panic("Invalid Logger/TagName.  Expected:" + tagPattern)
	}
	return id
}
