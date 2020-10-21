Dictionary
==========

## Purpose
* Provide an abstract `definition` struct for storing and handling dictionary 'words,' identifiers and related
  metadata with encryption/compression in-memory and at rest on disk.
* Provide a reusable reader/writer for the toolkit which will read/write a proprietary rich dictionary file format
  to/from disk.

## File Format
TBD
  
## Examples
### Reader
```go
package dictionary_example
import (
    "asymmetric-effort/asymmetric-toolkit/src/common/dictionary/reader"
    "os"
)

func doSomething(){
var reader DictionaryReader.Reader
file,err:=os.Create("myDictionary.txt")
if err != nil {
    panic(err)
}
defer func(){reader.Close()}()
read:=reader.Setup(file)
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
    "asymmetric-effort/asymmetric-toolkit/src/common/dictionary/writer"
    "os"
)

func doSomething(){
var writer DictionaryWriter.Writer
file,err:=os.Create("myDictionary.txt")
if err != nil {
    panic(err)
}
defer func(){writer.Close()}()
write:=writer.Setup(file)
write("my_word0")
write("my_word1")
write("my_word2")
write("my_word3")
}
```
