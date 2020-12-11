package dictionary

import "math"

func getPointerSize(p int64) uint32 {
	/*
		output should be 0...255
		output must be 32-bit to allow caller mask operations.
	*/
	switch {
	case p <= math.MaxUint8:
		return 1
	case p <= math.MaxUint16:
		return 2
	case p <= math.MaxUint32:
		return 4
	case uint(p) <= math.MaxUint64:
		return 8
	}
	panic("dictionary.GetPointerSize(): PointerSize out of bounds")
}
