Dictionary
==========

## Purpose
* Provide a single abstract 'definition' struct containing the dictionary 'word', a hash identifier and various
  metadata which can be used to prioritize words for use in various attack processes during a penetration test.
  
* Provide a reusable reader/writer for the toolkit dictionary file format which is encrypted on a record-by-record
  level using a symmetric AES passphrase and armored ciphertext which will not encounter binary storage issues in
  github.

## File Format
TBD
  
## Examples
### Reader
```go
package dictionary_example
import (
    "asymmetric-effort/asymmetric-toolkit/src/common/source/dictionary/reader"
    "os"
)

func doSomething(){
var r reader.Reader
file,err:=os.Create("myDictionary.txt")
if err != nil {
    panic(err)
}
defer func(){r.Close()}()
read:=r.Setup(file)
var eof bool
var data []string
eof,data = read()
read()
read()
read()
}
```

### Writer
```go
package dictionary_example
import (
    "asymmetric-effort/asymmetric-toolkit/src/common/source/dictionary/writer"
    "os"
)

func doSomething(){
var w writer.Writer
file,err:=os.Create("myDictionary.txt")
if err != nil {
    panic(err)
}
defer func(){w.Close()}()
write:=w.Setup(file)
write("my_word0")
write("my_word1")
write("my_word2")
write("my_word3")
}
```
