package logger

import "strconv"

func (o *LogSettings) SetInt(key string, value int) {
	(*o)[key] = strconv.Itoa(value)
}
