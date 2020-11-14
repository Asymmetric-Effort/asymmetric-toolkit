package logger

func (o *Logger) tagMerge(tags *[]TagId) []TagId {
	var mergedTags = make(TagTable, 1)
	for tag, _ := range o.tags.global {
		mergedTags[tag] = struct{}{}
	}
	if tags != nil {
		for _, tag := range *tags {
			mergedTags[tag] = struct{}{}
		}
	}
	var result []TagId
	for tag,_:=range mergedTags{
		result=append(result,tag)
	}
	return result
}
