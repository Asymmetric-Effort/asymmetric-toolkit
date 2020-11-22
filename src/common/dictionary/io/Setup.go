package io

import "crypto/sha256"

/*
	IO::Setup() - setup the dictionary compression and cryptography algorithm.
 */
func (o *IO) Setup(compressAlg Compression, passphrase string){
	if passphrase == "" {
		panic ("empty string is not allowed in dictionary encryption.")
	}
	o.Compress = compressAlg
	o.Crypto = sha256.Sum256([]byte(passphrase))
}