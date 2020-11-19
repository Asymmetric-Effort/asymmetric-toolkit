package dictionary

func (o *Dictionary) Compress(in *[]byte) (out *[]byte){
	return o.Config.CompressionAlg.Pack(in)
}