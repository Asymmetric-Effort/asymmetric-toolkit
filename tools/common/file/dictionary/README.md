Dictionary Reader/Writer
========================

## Purpose
* Provide a reusable reader/writer for dictionaries which we can use to extend more
  advanced dictionary functionality such as ranking and priority-based dictionary 
  sorts.
  
## Usage
### Reader
#### Setup(filePtr *os.File)
* Input: readable file handle (to the dictionary file).
* Output: reader function `func() (string, int int int...)`
#### Dictionary Reader Example
```go
package dictionary_example
import (
    "asymmetric-effort/asymmetric-toolkit/tools/common/file/dictionary"
    "os"
)

func doSomething(){
var reader dictionary.Reader
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
#### Setup(filePtr *os.File)
* Input: writable file handle (to the dictionary file).
* Output: writer function `func(s string)`

#### Dictionary Writer Example
```go
package dictionary_example
import (
    "asymmetric-effort/asymmetric-toolkit/tools/common/file/dictionary"
    "os"
)

func doSomething(){
var writer dictionary.Writer
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
