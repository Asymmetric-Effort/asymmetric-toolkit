package utils

import (
	"strconv"
)

func IsPort(s string) bool {
	// Is input a number 0-65535
	n, err := strconv.Atoi(s)
	return (err == nil) && (n >= 0) && (n <= 65535)
}