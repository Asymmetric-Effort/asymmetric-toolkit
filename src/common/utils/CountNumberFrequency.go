package utils

import "strconv"

func CountNumberFrequency(input int) *map[int]int {
	s := strconv.Itoa(input)

	result := make(map[int]int)
	for i := 0; i < 10; i++ { //initialize map
		result[i] = 0
	}
	for is, v := range *CountCharacterFrequency(&s) {
		i, err := strconv.Atoi(string(is))
		if err != nil {
			panic(err)
		}
		result[i] = v
	}
	return &result
}
