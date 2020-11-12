package logger

import "strconv"

func (o *LogSettings) GetInt(key string) int {
	v, err := strconv.Atoi((*o)[key])
	if err != nil {
		panic("LogSettings::SetInt expected integer.  Encountered non-integer input")
	}
	return v
}
