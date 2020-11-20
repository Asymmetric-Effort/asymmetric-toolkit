package definition

import (
	"bytes"
	"encoding/gob"
)

func (o *Descriptor) Deserialize(in *[]byte) {
	if o == nil {
		panic("nil Definition struct")
	}
	var data Descriptor
	dec := gob.NewDecoder(bytes.NewReader(*in))
	err := dec.Decode(&data)
	if err != nil {
		panic(err)
	}
	*o=data
}
