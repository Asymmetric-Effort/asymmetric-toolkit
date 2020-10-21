package definition

type Record struct {
	id      string
	//Note: file format puts a word length (int) here.
	word    string //Base64-encoded word
}
