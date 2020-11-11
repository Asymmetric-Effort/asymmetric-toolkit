package main

const (
	/*

	 */
	DefaultAttackDelay      int               = 1                    // Default number of seconds between attack payload executions.
	DefaultConcurrency      int               = 1                    // Default number of concurrent queries to run.
	DefaultAttackDepth      int               = 20                   // Default number of DNS subdomain levels to attack
	DefaultDnsDomain        string            = "localdomain"        // Default domain to be enumerated.
	DefaultDnsServer        string            = "udp://127.0.0.1:53" // Default DNS server.
	DefaultFilterPattern    string            = `.+`                 // Default source pattern.
	DefaultTimeout          int               = 60                   // Default number of seconds before connection timeout.
	DefaultMinWordSize      int               = 3                    // Default minimum character length of each word
	DefaultMaxWordSize      int               = 5                    // Default maximum character length of each word.
	DefaultOutput           string            = "output.txt"         // Default output path/file for reporting results.
	DefaultSource           string            = "sequence"           // Default word source for our enumeration attack.
	DefaultSourceDictionary string            = "my_dictionary.atk"  //Default path/file to a dictionary (we will have none).
	DefaultSourcePattern    string            = ".+"                 // Default regex pattern string to filter words to be used in an attack.
	//ResponseBufferSz int = 1048576 // Size of the response buffer for responses pending report processing.
	//SourceBufferSz   int = 1048576 // Size of the source buffer which feeds the payload (attack) function.
	DnsChars       string = "WMEFSABCDGHIJKLNOPQRTUVXYZwmefsabcdghijklnopqrtuvxyz0123456789._-"
	DnsRecordTypes        = "A,AAAA,MX,TXT,SOA,NS,CNAME"
)
