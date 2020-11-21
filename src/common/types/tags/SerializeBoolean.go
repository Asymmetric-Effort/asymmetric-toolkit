package tags

/*
	Tag::SerializeBoolean() is going to serialize the Tag struct to produce a byte slice.

		<tagCount:byte>
		<tagType:byte(1)><keyLength:byte><key:[]byte>
		<tagType:byte(1)><keyLength:byte><key:[]byte>
		.
		.
		.
		<tagType:byte(1)><keyLength:byte><key:[]byte>
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"bytes"
)

func (o *Boolean) SerializeBoolean() []byte {
	buf := bytes.Buffer{}
	buf.WriteByte(byte(len(*o))) // tagCount (uint32: 4 bytes)
	for key, value := range *o {
		buf.WriteByte(byte(BooleanTag))       // BooleanTag key.
		buf.WriteByte(byte(len(key)))         // Length of key
		buf.WriteByte(misc.BoolToByte(value)) // Boolean value (byte).
		buf.Write([]byte(key))                // key (bytes)
	}
	//fmt.Println("String::Serialize(): start d=", buf.Bytes())
	return buf.Bytes()
}
