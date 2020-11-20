package headers

import "asymmetric-effort/asymmetric-toolkit/src/common/dictionary"

/*
	Descriptor - This is the dictionary file header.
*/
type Descriptor struct {
	FormatVersion dictionary.Version // uint8  : This is the version for the file format.
	ScoreVersion  dictionary.Version // uint8  : This is the version of the score.
	WordCount     uint32             // uint32 : Number of definitions in the file.
}
