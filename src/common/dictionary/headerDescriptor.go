package dictionary

type HeaderDescriptor struct {
	FormatVersion Version // uint8
	ScoreVersion  Version // uint8
	WordCount     uint32  //Number of definitions in the file
}

func (o *Dictionary) CreateHeader() {
	o.Content.Header.FormatVersion = o.Config.FormatVersion
	o.Content.Header.ScoreVersion = o.Config.ScoreVersion
	o.Content.Header.WordCount = 0 // We will populate this later.
}
