package tags

/*
	Tag::SerializeString() is going to serialize the Tag struct to produce a byte slice.

		<tagCount:byte>
		<tagType:byte(1)><keyLength:byte><key:[]byte><valueLen:byte><value:[]byte>
		<tagType:byte(1)><keyLength:byte><key:[]byte><valueLen:byte><value:[]byte>
		.
		.
		.
		<tagType:byte(1)><keyLength:byte><key:[]byte><valueLen:byte><value:[]byte>
*/
import (
	"bytes"
)

func (o *String) SerializeString() []byte {
	mutex.Lock()
	defer mutex.Unlock()

	buf := bytes.Buffer{}
	buf.WriteByte(byte(len(*o))) // tagCount (uint8: 1 byte)
	for key, value := range *o {
		buf.WriteByte(byte(StringTag))  // StringTag key.
		buf.WriteByte(byte(len(key)))   // Length of key (Byte)
		buf.WriteByte(byte(len(value))) // Length of value (Byte)
		buf.Write([]byte(key))          // key (bytes)
		buf.Write([]byte(value))        // value (bytes)
	}
	//fmt.Println("String::Serialize(): start d=", buf.Bytes())
	return buf.Bytes()
}
