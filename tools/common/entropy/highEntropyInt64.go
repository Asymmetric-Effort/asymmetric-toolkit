package entropy

func HighEntropyInt64(s int64) bool {
	return GetShannonsInt64(s) >= HighEntropyThreshold
}
