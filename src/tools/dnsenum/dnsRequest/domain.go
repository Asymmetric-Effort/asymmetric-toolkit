package dnsRequest

import (
	"fmt"
	"os"
)

func (o *Request) Domain(d *string) {
	if d == nil {
		fmt.Println("Errorf(Request::Domain()): null Domain encountered.")
		os.Exit(1)
	}
	if *d == "" {
		fmt.Println("Errorf(Request::Domain()): empty Domain encountered.")
		os.Exit(1)
	}
	o.domain = d
}
