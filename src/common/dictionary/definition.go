package dictionary

type Definition struct {
	Id      uint32
	Word    string
	Score   int
	Created int64          // Unix nano timestamp when the word is identified/created.
	LastHit int64          // Unix nano timestamp when the word is identified/created.
	Tags    DefinitionTags // Tags used to identify definition attributes.
	Hits    int            //Count of hits
	Miss    int            //Count of misses
}
