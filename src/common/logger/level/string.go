package level

import (
	"fmt"
	"strings"
)

func (o *Level) String() (s string) {
	list := strings.Split(Strings, ",")
	i := int(o.Get())
	if i < len(list) {
		return list[i]
	} else {
		panic(fmt.Sprintf("Programming error: Invalid Log level: %d", *o))
	}
}
