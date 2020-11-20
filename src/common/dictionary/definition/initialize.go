package definition

import "asymmetric-effort/asymmetric-toolkit/src/common/types/tags"

func (o *Descriptor) Initialize() {
	/*
		definition.Descriptor initializes the struct with initial state for new definitions.
	*/
	if o == nil {
		panic("nil Definition struct")
	}
	o.Id = 0               //uint32
	o.Word = ""            //string
	o.Score = 0            //int
	o.Created = 0          //int64          // Unix nano timestamp when the word is identified/created.
	o.LastHit = 0          //int64          // Unix nano timestamp when the word is identified/created.
	o.Tags = tags.String{} //DefinitionTags // Tags used to identify definition attributes.
	o.Hits = 0             //int              //Count of hits
	o.Miss = 0             //int              //Count of misses
}
