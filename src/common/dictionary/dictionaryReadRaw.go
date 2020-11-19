package dictionary

func (o *Dictionary) ReadRaw() *[]byte {
	return o.Decompress(o.Decrypt(o.ReadBytes(o.ReadBytes(recordLengthSize))))
}
