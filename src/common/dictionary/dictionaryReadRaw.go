package dictionary

func (o *Descriptor) ReadRaw() *[]byte {
	return o.Decompress(o.Decrypt(o.ReadBytes(o.ReadBytes(recordLengthSize))))
}
