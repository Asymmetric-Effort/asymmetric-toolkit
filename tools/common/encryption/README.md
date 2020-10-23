Encryption
==========

## Purpose
Provide abstract encryption functions and tooling to simplify maintenance of the encryption pieces and
make development easier for various parts of the Asymmetric Toolkit.

## Tools

### Key (type)
* Provides an encryption key type for abstract use beyond the standard [32]byte.
* Facilitates future transition to other hash algorithms over time.

#### Methods

##### Set(passphrase *string)
Takes string pointer and produces internal 32-byte sha256 hash state (key)

##### String() string
Returns a hex string for the current internal state.

##### Length() int
Returns the current hash length.  Usage is recommended as it allows a more
graceful move to SHA512 or SHA3 algorithms in the future.

##### Equal(hash *[32]byte) bool
Compares the internal state to a given 32-byte array and returns boolean result.

### Encrypt() *string
Given a string, this method will encrypt the cleartext string and return a hexadecimal
string containing the ciphertext.

### Decrypt() *string
Given a hexadecimal string containing ciphertext, this function decrypts the signal
and returns the cleartext string.