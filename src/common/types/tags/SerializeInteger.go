package tags

/*
	Tag::SerializeInteger() is going to serialize the Tag struct to produce a byte slice.

		<tagCount:byte>
		<tagType:byte(1)><keyLength:byte><key:[]byte><value:uint32>
		<tagType:byte(1)><keyLength:byte><key:[]byte><value:uint32>
		.
		.
		.
		<tagType:byte(1)><keyLength:byte><key:[]byte><value:uint32>
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"bytes"
)

func (o *Integer) SerializeInteger() []byte {
	mutex.Lock()
	defer mutex.Unlock()

	buf := bytes.Buffer{}
	buf.WriteByte(byte(len(*o))) // tagCount (uint8: 1 byte)
	for key, value := range *o {
		buf.WriteByte(byte(IntegerTag))             // IntegerTag key.
		buf.WriteByte(byte(len(key)))               // Length of key (Byte)
		buf.Write([]byte(key))                      // key (bytes)
		buf.Write(misc.Uint32ToByte(uint32(value))) // value (bytes)
	}
	//fmt.Println("String::Serialize(): start d=", buf.Bytes())
	return buf.Bytes()
}
