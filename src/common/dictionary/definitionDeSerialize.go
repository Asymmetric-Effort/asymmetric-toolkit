package dictionary

func (o *Definition) Deserialize(data *[]byte) {
	if o == nil {
		panic("nil Definition struct")
	}
	o.Id = 0                    //uint32
	o.Word = ""                 //string
	o.Score = 0                 //int
	o.Created = 0               //int64          // Unix nano timestamp when the word is identified/created.
	o.LastHit = 0               //int64          // Unix nano timestamp when the word is identified/created.
	o.Tags = []DefinitionTags{} //DefinitionTags // Tags used to identify definition attributes.
	o.Hits = 0                  //int              //Count of hits
	o.Miss = 0                  //int              //Count of misses
}
