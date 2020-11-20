package dictionary

func (o *Descriptor) WriteRaw(data *[]byte) {
	o.WriteBytes(o.Encrypt(o.Compress(data)))

	//return o.Decompress(o.Decrypt(o.ReadBytes(recordLengthSize)))
}
