package misc

func CountCharacterFrequency(s *string) *map[rune]float64 {
	result := make(map[rune]float64)
	for _, i := range *s {
		result[i]++
	}
	return &result
}