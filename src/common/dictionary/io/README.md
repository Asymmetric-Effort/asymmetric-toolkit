Dictionary I/O
==============

## Purpose
* To provide compression and cryptography functionality for the Asymmetric-Toolkit
  dictionary files.

## Example
The top-level uses the `IO` object to `Encode()` or `Decode()` information:
```go
package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/dictionary/io"
    "fmt"
)

func main(){
var i io.IO 
    i.Setup(io.Gzip, "myPassphrase")
    in:=[]byte("My source text.")
    out:=i.Encode(&in)
    verify:=*i.Decode(out)
    fmt.Println(verify)
}
```
These two methods will wrap the compression and encryption processes.
