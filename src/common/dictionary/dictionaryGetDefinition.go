package dictionary

/*
	Descriptor::GetDefinition() read file for one definition object encrypted and compressed.

	<length:int><byte array of encrypted, compressed data...>
*/
func (o *Descriptor) GetDefinition() (def Definition) {
	raw := o.ReadRaw()
	var definition Definition
	definition.Deserialize(raw)
	//Get the next definition.
	return definition
}
