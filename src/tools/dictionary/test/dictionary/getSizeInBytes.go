package dictionary

import (
	"encoding/binary"
	"fmt"
	"math"
)

func (o *Word) GetSizeInBytes() []byte {
	maxWordLen := int(math.Pow(2.0, float64(wordSizeLength*8)))
	sz := len(o.Word)
	if sz >= maxWordLen {
		panic(fmt.Sprintf("Word::GetSizeInBytes() "+
			"Error: word length exceeds %d (%d)", maxWordLen, sz))
	}
	wordBytes := make([]byte, wordSizeLength)
	binary.BigEndian.PutUint64(wordBytes, uint64(len(o.Word)))
	return wordBytes
}
