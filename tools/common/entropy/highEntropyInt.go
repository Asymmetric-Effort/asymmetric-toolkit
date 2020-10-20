package entropy

func HighEntropyInt(s int) bool {
	return GetShannonsIntArray(s) >= HighEntropyThreshold
}
