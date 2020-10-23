package entropy

func HighEntropy(inputString string) bool {
	return GetShannons(inputString) >= HighEntropyThreshold
}
