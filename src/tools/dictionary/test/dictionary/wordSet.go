package dictionary

import "asymmetric-effort/asymmetric-toolkit/src/common/errors"

func (o *Word) Set(word *string, lhs int64, rhs int64) {
	const maxWordLength = 65535
	length := len(*word)
	errors.Assert(length <= maxWordLength, "Words cannot be more than 65535 bytes")
	errors.Assert(lhs >= 0, "Expect lhs >= 0")
	errors.Assert(rhs >= 0, "Expect rhs >= 0")
	//o.WordSz = length
	o.Word = *word
	o.Lhs = lhs
	o.Rhs = rhs
}