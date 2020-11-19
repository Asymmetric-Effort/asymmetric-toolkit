Dictionary Reader/Writer
========================

## Purpose
Provide a reusable reader/writer for the toolkit dictionary file format which is 
encrypted on a record-by-record level using a symmetric AES passphrase and armored
ciphertext which will not encounter binary storage issues in github.


## Usage
```go
package main

import (
    "asymmetric-effort/asymmetric-toolkit/src/common/dictionary"
)

func main(){
    var source dictionary.Dictionary
    var target dictionary.Dictionary
    source.Setup(&dictionary.Configuration{
        FileName: "myDict.atd",
        Overwrite: true,
        FormatVersion: 0,
        ScoreVersion: 0,
        Passphrase: []("passphrase"),
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

    sourceHeader:=source.GetHeader()

    for definition:=range source.GetDefinition() {
        if err != nil {
            panic(err)
        }
    	target.AddDefinition(&definition)
    }
    target.SetHeader(&sourceHeader)
}
```
