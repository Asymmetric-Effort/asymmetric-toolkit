package types

import (
	"fmt"
	"strings"
)

type List struct {
	data []string
}

func (o *List) Set(s string) {
	// Consume comma-delimited list as string.
	// Split by comma delimiter and store the list.
	o.data = append(o.data, strings.Split(s, ",")...)
}

func (o *List) Count() int {
	return len(o.data)
}

func (o *List) Get(i int) string {
	if i < len(o.data) {
		return o.data[i]
	} else {
		panic(fmt.Sprintf("List.Get() index (%d) out of bounds", i))
	}
}

func (o *List) String() string {
	return strings.Join(o.data, ",")
}
