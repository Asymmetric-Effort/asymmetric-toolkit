package dictionary

type DefinitionTags struct {
	tag []string
}

func (o *DefinitionTags) Find(s string) bool {
	for _, value := range o.tag {
		if s == value {
			return true
		}
	}
	return false
}
