Dictionary Reader/Writer
========================

## Purpose
Provide a reusable reader/writer for the toolkit dictionary file format which is 
encrypted on a record-by-record level using a symmetric AES passphrase and armored
ciphertext which will not encounter binary storage issues in github.


## Usage
The following is a comprehensive example of all features of the reader/writer dictionary:
```go
package main

import (
    "asymmetric-effort/asymmetric-toolkit/src/common/dictionary"
    "github.com/google/uuid"
    "time"
)

func main(){
    var source dictionary.Dictionary
    var target dictionary.Dictionary
    source.Setup(&dictionary.Configuration{
        FileName: "myDict.atd",
        Overwrite: true,
        FormatVersion: 0,
        ScoreVersion: 0,
        Passphrase: []byte("passphrase"),
        CompressionAlg: dictionary.Gzip,
    })

    target.Setup(&dictionary.Configuration{
        FileName: "myDict.atd",
        Overwrite: true,
        FormatVersion: 0,
        ScoreVersion: 0,
        Passphrase: []byte("passphrase"),
        CompressionAlg: dictionary.Gzip,
    })

    sourceHeader,err:=source.GetHeader()
    if err != nil {
        panic(err)
    }

    for definition:=range source.GetDefinition() {
    	target.AddDefinition(&definition)
    }
    target.SetHeader(&sourceHeader)

    newDefinitionId,err:=uuid.NewUUID()
    if err!= nil{
        panic(err)
    }
    
    var tags= dictionary.DefinitionTags{
        "my_tag":"my_value",
    }

    target.AddDefinition(&dictionary.Definition{
        Id      :newDefinitionId, //uuid.UUID
        Word    :"",                    // string
        Score   :0,                     // int
        Created :time.Now().UnixNano(), // Unix nano timestamp when the word is identified/created.
        LastHit :0,                     // Unix nano timestamp when the word is identified/created.
        Tags    :tags,                  // Tags used to identify definition attributes.
        Hits    :0,                     // Count of hits
        Miss    :0,                     // Count of misses
    })


}
```
