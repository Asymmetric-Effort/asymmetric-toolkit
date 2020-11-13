package logger

/*
	Level::String() converts the numeric log level to its string equivalent.
*/
import (
	"fmt"
	"strings"
)

func (o *Level) String() (s string) {
	list := strings.Split(LevelStrings, ",")
	i := int(o.Get())
	if i < len(list) {
		s = strings.ToLower(list[i])
	} else {
		panic(fmt.Sprintf("Programming error: Invalid Log level: %d", *o))
	}
	return
}
