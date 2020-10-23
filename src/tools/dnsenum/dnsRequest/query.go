package dnsRequest

import (
	"fmt"
)

func (o *Request) Query(record *string) {
	fmt.Print("record:", record)
}
