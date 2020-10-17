package entropy

import "math"

func CountCharacterFrequency(s *string) *map[rune]float64 {
	result := make(map[rune]float64)
	for _, i := range *s {
		result[i]++
	}
	return &result
}

func GetShannons(inputString string) int {
	//See https://en.wikipedia.org/wiki/Entropy_(information_theory)
	//See https://en.wikipedia.org/wiki/Shannon_(unit)
	//See https://en.wikipedia.org/wiki/ISO/IEC_80000
	//See https://en.wiktionary.org/wiki/Shannon_entropy

	var sum float64
	var strLen int = len(inputString)
	for _, v := range *CountCharacterFrequency(&inputString) {
		var f float64 = v / float64(strLen)
		sum += f * math.Log2(f)
	}
	return int(math.Ceil(sum*-1)) * strLen
}

func HighEntropy(inputString string) bool {
	return GetShannons(inputString) >= 128
}
