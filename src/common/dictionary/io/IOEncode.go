package io
/*
	io.Encode() - this is the top-level encoder for the dictionary which encrypts the
   				  compressed signal.
 */
func (o *IO) Encode(in *[]byte) (out *[]byte) {
	return o.Crypto.Encrypt(o.Compress.Pack(in))
}
