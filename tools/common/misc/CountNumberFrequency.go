package misc

import "strconv"

func CountNumberFrequency(input int) *map[int]int {
	s := strconv.Itoa(input)

	result := make(map[int]int)
	for is, v := range *CountCharacterFrequency(&s) {
		i,err:=strconv.Atoi(string(is))
		if err != nil {
			panic (err)
		}
		result[i] = v
	}
	return &result
}
