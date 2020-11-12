package logger

import "strconv"

func (o *LogSettings) SetBool(key string, value bool) {
	(*o)[key] = strconv.FormatBool(value)
}
