package definition

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types/tags"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestDescriptor_Serialize(t *testing.T) {
	const (
		uint32Sz = 4
		uint64Sz = 8
	)

	var o Descriptor
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	errors.Assert(o.Id == u, "pre-check: expect 0")
	errors.Assert(o.Word == "", "pre-check: expect empty string")
	errors.Assert(o.Score == 0, "pre-check: expect 0")
	errors.Assert(o.Created == 0, "pre-check: expect 0")
	errors.Assert(o.LastHit == 0, "pre-check: expect 0")
	errors.Assert(&o.Tags != nil, "pre-check: expect non-nil pointer")
	errors.Assert(o.Hits == 0, "pre-check: expect 0")
	errors.Assert(o.Miss == 0, "pre-check: expect 0")
	fmt.Println("Pre-check [all]: OK")

	o.Word = "test"
	tagLen := func(t tags.String) int {
		var n int
		for _, v := range o.Tags {
			n += len(v)
		}
		return n
	}
	WordSz := len(o.Word)
	TagSz := tagLen(o.Tags)
	srcLen := uint32Sz + uint32Sz + WordSz + uint32Sz + uint64Sz + uint64Sz + TagSz + uint32Sz + uint32Sz

	data := o.Serialize() // Serialize struct -> []byte
	fmt.Println("data:", len(data), "/", srcLen, " : ", data)
}
