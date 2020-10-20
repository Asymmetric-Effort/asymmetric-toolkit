package misc

import "strconv"

func CountNumberFrequency64(input int64) *map[int]int {
	s := strconv.FormatInt(input, 10)

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
