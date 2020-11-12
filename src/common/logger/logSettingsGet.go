package logger

import (
	"fmt"
)


func (o *LogSettings) Get(key string) string {
	if value, ok := (*o)[key]; ok {
		return value
	}
	panic(fmt.Sprintf("LogSettings::Get() requested unknown key (%s)", key))
}