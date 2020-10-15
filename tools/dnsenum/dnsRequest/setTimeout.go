package dnsRequest

import (
	"fmt"
	"os"
	"time"
)

func (o *Request) SetTimeout(t *time.Duration) {
	if t == nil {
		fmt.Println("Error(Request::setTimeout()): null timeout encountered.")
		os.Exit(1)
	}
	o.timeout = t
}
