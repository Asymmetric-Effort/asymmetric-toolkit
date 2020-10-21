package dnsRequest

import (
	"fmt"
	"os"
)

func (o *Request) SetConcurrency(c int) {
	if c < 0 {
		fmt.Println("Errorf(Request::setConcurrency()): must be greater than 0.")
		os.Exit(1)
	}
	o.concurrency = c
}

