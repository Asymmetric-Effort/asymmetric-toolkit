package tags

/*
	Tag::SerializeBoolean() is going to serialize the Tag struct to produce a byte slice.

		<tagCount:uint32>
		<tagType:byte(1)><keyLength:uint32><key:[]byte>
		<tagType:byte(1)><keyLength:uint32><key:[]byte>
		.
		.
		.
		<tagType:byte(1)><keyLength:uint32><key:[]byte>
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"bytes"
)

func (o *Boolean) SerializeBoolean() []byte {
	buf := bytes.Buffer{}
	buf.Write(misc.Uint32ToByte(uint32(len(*o)))) // tagCount (uint32: 4 bytes)
	for key, value := range *o {
		buf.WriteByte(byte(BooleanTag))                // BooleanTag key.
		buf.Write(misc.Uint32ToByte(uint32(len(key)))) // Length of key
		buf.WriteByte(misc.BoolToByte(value))          // Boolean value (byte).
		buf.Write([]byte(key))                         // key (bytes)
	}
	//fmt.Println("String::Serialize(): start d=", buf.Bytes())
	return buf.Bytes()
}
