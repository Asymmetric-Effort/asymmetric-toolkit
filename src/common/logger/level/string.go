package level

import (
	"fmt"
	"strings"
)

func (o *LogLevel) String() (s string) {
	list := strings.Split(LevelStrings, ",")
	i := int(o.Get())
	if i < len(list) {
		return list[i]
	} else {
		panic(fmt.Sprintf("Programming error: Invalid Log level: %d", *o))
	}
}
