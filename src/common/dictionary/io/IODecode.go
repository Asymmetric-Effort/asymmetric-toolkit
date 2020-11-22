package io
/*
	io.Decode() - this is the top-level decoder for the dictionary which decompresses the
				  decrypted signal.
 */
func (o *IO) Decode(in *[]byte) (out *[]byte) {
	return o.Compress.Unpack(o.Crypto.Decrypt(in))
}
