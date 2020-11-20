package definition

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"bytes"
)

func (o *Descriptor) Serialize() []byte {
	buf := bytes.Buffer{}

	buf.Write(misc.UuidSerialize(o.Id))               // definition ID (16-bytes)
	buf.Write(misc.Uint32ToByte(uint32(len(o.Word)))) // definition Word length
	buf.Write([]byte(o.Word))                         // definition Word string
	buf.Write(misc.Uint64ToByte(o.Created))           // definition Created
	buf.Write(misc.Uint64ToByte(o.LastHit))           // definition LastHit
	buf.Write(o.Tags.Serialize())                     // definition tagLength
	buf.Write(misc.Uint32ToByte(o.Hits))              // definition Hits
	buf.Write(misc.Uint32ToByte(o.Miss))              // definition Miss

	return buf.Bytes()
}
