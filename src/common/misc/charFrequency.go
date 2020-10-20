package misc

func CountCharacterFrequency(s *string) *map[rune]int {
	result := make(map[rune]int)
	for _, i := range *s {
		result[i]++
	}
	return &result
}
