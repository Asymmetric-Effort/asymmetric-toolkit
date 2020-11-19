Dictionary Reader/Writer
========================

## Purpose
Provide a reusable reader/writer for the toolkit dictionary file format which is 
encrypted on a record-by-record level using a symmetric AES passphrase and armored
ciphertext which will not encounter binary storage issues in github.


# Theory of Operation
```go 
type Dictionary struct{
    Initialize func()
    OpenRead func(fileName string)
    OpenWrite func(fileName string)
    CloseFile func()
    
    LoadHeader func() *Header
    LoadDefinition func() *Definition //First defnition reference.
    Next func() *Definition
    
    handle *os.File
   
    content struct {
        header struct {
            formatVersion uint8
            scoreVersion  uint8
            reserved      uint16
            wordCount     uint32
        }
        definitions []struct{
            Id      uint32
            Word    string
            Score   int
            Created int64          // Unix nano timestamp when the word is identified/created.
            LastHit int64          // Unix nano timestamp when the word is identified/created.
            tags    DefinitionTags // Tags used to identify definition attributes.
            hits    int            //Count of hits
            miss    int            //Count of misses
        }
    }
}


func (o *Dictionary) Setup(){

    var fileHandle *os.File

    o.OpenRead:=func(fileName string){
        //Open the file
    }
    o.OpenWrite:=func(fileName string){
        //Open the file.
    }
}

```
