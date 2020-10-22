package utils

import "strconv"

func IntLen64(i int64) int64 {
	s := strconv.Itoa(int(i))
	return int64(len(s))
}
