package misc

import "github.com/google/uuid"

func UuidSerialize(id uuid.UUID) []byte {
	uuidBytes, err := id.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return uuidBytes
}