package dnsRequest

import (
	"time"
)

type Request struct {
	concurrency    int
	//activeRequests int

	timeout *time.Duration
	domain  *string
	dnsType *string
}
