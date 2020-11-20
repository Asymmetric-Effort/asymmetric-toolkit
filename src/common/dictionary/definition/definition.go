package definition

/*
	Definition - A definition is a structure with a word, an identifier and various statistics used to prioritize
				 words used in attacks.  Each definition is serialized/deserialized as it is read from or written
				 to disk.
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/types/tags"
)

type Descriptor struct {
	Id      uint32
	Word    string
	Score   int
	Created int64       // Unix nano timestamp when the word is identified/created.
	LastHit int64       // Unix nano timestamp when the word is identified/created.
	Tags    tags.String // Tags used to identify definition attributes.
	Hits    int         //Count of hits
	Miss    int         //Count of misses
}

