package dictionary

import (
	"fmt"
	"math"
)

func getPointerSize(p int64) (result uint32) {
	/*
		output should be 0...255
		output must be 32-bit to allow caller mask operations.
	*/
	switch {
	case p <= math.MaxUint8:
		result= 1
	case p <= math.MaxUint16:
		result= 2
	case p <= math.MaxUint32:
		result=4
	case uint(p) <= math.MaxUint64:
		result=8
	default:
		panic("dictionary.GetPointerSize(): PointerSize out of bounds")
	}
	fmt.Printf("\t\tdictionary.getPointerSize(): result:%d\n",result)
	return result
}
