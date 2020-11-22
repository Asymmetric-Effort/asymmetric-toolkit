package definition

/*
	definition.Descriptor::Serialize() - Serialize a dictionary definition structure
*/
import (
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"bytes"
)

func (o *Descriptor) Serialize() []byte {
	buf := bytes.Buffer{}
	buf.Write(misc.UuidSerialize(o.Id))               // ...ID (16-bytes)
	buf.Write(misc.Uint32ToByte(uint32(len(o.Word)))) // ...Word length (32 bits)
	buf.Write(misc.Uint64ToByte(o.Created))           // ...Created (64-bit)
	buf.Write(misc.Uint64ToByte(o.LastHit))           // ...LastHit (64-bit)
	buf.Write(misc.Uint32ToByte(o.Hits))              // ...Hits (32-bit)
	buf.Write(misc.Uint32ToByte(o.Miss))              // ...Miss (32-bit)
	buf.Write(o.Tags.SerializeString())               // ...tag (variable length)
	buf.Write([]byte(o.Word))                         // ...Word string (variable length)
	return buf.Bytes()                                // return the final buffer
}
