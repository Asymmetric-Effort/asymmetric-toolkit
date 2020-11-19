package dictionary

/*
	Dictionary::GetDefinition() read file for one definition object encrypted and compressed.

	<length:int><byte array of encrypted, compressed data...>
*/
func (o *Dictionary) GetDefinition() (def *Definition) {
	raw := o.ReadRaw()
	var definition Definition
	definition.Deserialize(raw)
	//Get the next definition.
	return &definition
}
