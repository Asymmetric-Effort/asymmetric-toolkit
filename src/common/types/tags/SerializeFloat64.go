package tags

/*
	Tag::SerializeFloat64() is going to serialize the Tag struct to produce a byte slice.

		<tagCount:byte>
		<tagType:byte(1)><keyLength:byte><key:[]byte><value:uint64>
		<tagType:byte(1)><keyLength:byte><key:[]byte><value:uint64>
		.
		.
		.
		<tagType:byte(1)><keyLength:byte><key:[]byte><value:uint64>
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"bytes"
	"math"
)

func (o *Float64) SerializeFloat64() []byte {
	mutex.Lock()
	defer mutex.Unlock()

	buf := bytes.Buffer{}
	buf.WriteByte(byte(len(*o))) // tagCount (uint8: 1 byte)
	for key, value := range *o {
		buf.WriteByte(byte(Float64Tag))                       // Float64 Tag type.
		buf.WriteByte(byte(len(key)))                         // Length of key (Byte)
		buf.Write(misc.Uint32ToByte(uint32(len(key))))        // Length of key
		buf.Write(misc.Uint64ToByte(math.Float64bits(value))) // binary Float64 value (byte). (IEEE 754 format)
	}
	//fmt.Println("String::Serialize(): start d=", buf.Bytes())
	return buf.Bytes()
}
