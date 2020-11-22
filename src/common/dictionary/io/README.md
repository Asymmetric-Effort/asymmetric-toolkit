Dictionary I/O
==============

## Purpose
* To provide compression and cryptography functionality for the Asymmetric-Toolkit
  dictionary files.

## Example
The top-level uses the `IO` object to `Encode()` or `Decode()` information:
```go
package main

func main(){
var io IO 
    io.Setup(io.Gzip, "mypassphrase")
    in:="My source text."
    out:=IO.encode(&in)
    verify:=*IO.decode(&out)
}
```
These two methods will wrap the compression and encryption processes.
