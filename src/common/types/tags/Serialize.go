package tags

/*
	Tag::Serialize() is going to serialize the Tag struct to produce a byte slice.

		<tagCount:byte>
		<tagType:byte(0)><keyLength:byte><key:[]byte>
		<tagType:byte(0)><keyLength:byte><key:[]byte>
		.
		.
		.
		<tagType:byte(0)><keyLength:byte><key:[]byte>
*/
import (
	"bytes"
)

func (o *Tag) Serialize() []byte {
	mutex.Lock()
	defer mutex.Unlock()

	buf := bytes.Buffer{}

	buf.WriteByte(byte(len(*o))) // tagCount (uint8: 1 byte)
	for key, _ := range *o {
		buf.WriteByte(byte(BasicTag)) // BasicTag is only a key (no value).
		buf.WriteByte(byte(len(key))) // Length of key (Byte)
		buf.Write([]byte(key))        // key (bytes)

	}

	//fmt.Println("String::Serialize(): start d=", buf.Bytes())

	return buf.Bytes()
}
