package dictionary

func (o *Dictionary) Decompress(in *[]byte) (out *[]byte){
	return o.Config.CompressionAlg.Unpack(in)
}