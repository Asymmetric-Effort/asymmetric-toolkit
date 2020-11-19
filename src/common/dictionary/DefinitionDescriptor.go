package dictionary

type DefinitionDescriptor struct {
	Id      uint32
	Word    string
	Score   int
	Created int64          // Unix nano timestamp when the word is identified/created.
	LastHit int64          // Unix nano timestamp when the word is identified/created.
	tags    DefinitionTags // Tags used to identify definition attributes.
	hits    int            //Count of hits
	miss    int            //Count of misses
}
