package cli

import "asymmetric-effort/asymmetric-toolkit/src/common/types"

const (
	DefaultConcurrency    types.PositiveInteger = 1  //default number of concurrent queries to run.
	DefaultDepth          types.PositiveInteger = 1  //default number of DNS subdomain levels to attack
	MaxDepth              types.PositiveInteger = 20 //Maximum number of DNS subdomain levels to attack
	DefaultTimeout        types.PositiveInteger = 60 //Default number of seconds before connection timeout.
	DefaultFilterPattern  string                = `.+`
	DefaultDNSRecordTypes string                = "A,AAAA,MX,CNAME,NS,TXT,SOA,SRV"
	DefaultWordSize       types.PositiveInteger = 5 //For a given sequence or random string, this is the max length of the 'word'
	DNSChars              string                = "WMEFSABCDGHIJKLNOPQRTUVXYZwmefsabcdghijklnopqrtuvxyz0123456789._-"
	SourceBufferSz        int                   = 1048576 //Size of the source buffer which feeds the payload (attack) function.
	ResponseBufferSz      int                   = 1048576 //Size of the response buffer for responses pending report processing.

)
