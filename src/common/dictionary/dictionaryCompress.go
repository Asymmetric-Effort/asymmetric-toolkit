package dictionary

func (o *Descriptor) Compress(in *[]byte) (out *[]byte){
	return o.Config.CompressionAlg.Pack(in)
}