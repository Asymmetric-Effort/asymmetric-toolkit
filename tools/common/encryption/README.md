Encryption
==========

## Purpose
Provide abstract encryption functions and tooling to simplify maintenance of the encryption pieces and
make development easier for various parts of the Asymmetric Toolkit.

## Tools

### CreateKey()

##### Input: 
* string (passphrase)
##### Output: 
* []byte (32-byte AES encryption key)


### Encrypt()

##### Input: 
* cleartext (string pointer): message to be encrypted
* key (byte slice): encryption key (32 bytes)
##### Output:
* Base64 encoded ciphertext (string)
 
 
### Decrypt()

##### Input: 
* cypherText (string pointer): base64-encoded message to be decrypted
* key (byte slice): encryption key (32 bytes)
##### Output:
* cleartext (string pointer).
