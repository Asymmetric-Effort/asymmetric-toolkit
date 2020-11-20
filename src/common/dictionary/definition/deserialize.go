package definition

import "asymmetric-effort/asymmetric-toolkit/src/common/types/tags"

func (o *Descriptor) Deserialize(data *[]byte) {
	if o == nil {
		panic("nil Definition struct")
	}
	o.Id = 0               // Definition (word) UUID identifier
	o.Word = ""            // Definition (word) string
	o.Score = 0            // int score
	o.Created = 0          // Unix nano timestamp when the word is identified/created.
	o.LastHit = 0          // Unix nano timestamp when the word is identified/created.
	o.Tags = tags.String{} // Tags used to identify definition attributes.
	o.Hits = 0             // Count of hits
	o.Miss = 0             // Count of misses
}
