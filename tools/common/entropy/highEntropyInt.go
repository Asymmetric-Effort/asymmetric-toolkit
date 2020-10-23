package entropy

func HighEntropyInt(s int) bool {
	return GetShannonsInt(s) >= HighEntropyThreshold
}
