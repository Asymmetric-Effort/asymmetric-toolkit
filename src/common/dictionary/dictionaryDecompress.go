package dictionary

func (o *Descriptor) Decompress(in *[]byte) (out *[]byte){
	return o.Config.CompressionAlg.Unpack(in)
}