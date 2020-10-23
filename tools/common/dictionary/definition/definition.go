package dictionaryDefinition

type Record struct {
	id      string
	//ToDo: Add digested stats here
	//Note: file format puts a word length (int) here.
	word    string //Base64-encoded word
}
