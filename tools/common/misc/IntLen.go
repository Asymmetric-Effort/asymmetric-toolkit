package misc

import "strconv"

func IntLen(i int) int{
	s:=strconv.Itoa(i)
	return len(s)
}