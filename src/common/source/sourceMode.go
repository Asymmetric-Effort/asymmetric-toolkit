package source

type Mode uint8

const (
	Sequence   Mode = 0 // Produce a linear sequence of a given keyspace.
	Random     Mode = 1 // Produce random sequences of a given keyspace.
	Dictionary Mode = 2 // Produce words from a dictionary file.

	strModeSequence = "sequence"
	strModeRandom = "random"
	strModeDictionary = "dictionary"
)
