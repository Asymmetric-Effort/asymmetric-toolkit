package definition

/*
	Definition - A definition is a structure with a word, an identifier and various statistics used to prioritize
				 words used in attacks.  Each definition is serialized/deserialized as it is read from or written
				 to disk.
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/types/tags"
	"github.com/google/uuid"
)

type Descriptor struct {
	Id      uuid.UUID   // Definition Identifier
	Score   uint        // weighted score.
	Created uint64      // Unix nano timestamp when the word is identified/created.
	LastHit uint64      // Unix nano timestamp when the word is identified/created.
	Hits    uint32      //Count of hits
	Miss    uint32      //Count of misses
	Word    string      // string word (cleartext)
	Tags    tags.String // Tags used to identify definition attributes.
}
