package cli

/*
	Common flag identifiers intended to promote standardization across the asymmetric toolkit.  For all
	reusable commandline flags, we reserve 0-1000.  For program-specific flags not shared across the tool suite,
	we reserve numbers > 1000.

	NOTE: OTHER THAN <=1000 BEING RESERVED TO THIS FILE, NO NUMERIC VALUE SHOULD BE
	      GIVEN ANY RELATIVE MEANING OVER ANY OTHER VALUE.
*/
type ArgumentFlag int

const ( //Reserved 0...1000 for core use.  1000+ are program-defined.
	noFlag             ArgumentFlag = 0 // noop
	FlagHelp           ArgumentFlag = 1 // Displays usage
	FlagVersion        ArgumentFlag = 2 // Displays version string
	FlagDebug          ArgumentFlag = 3 // Indicates debug logging.
	FlagLogLevel       ArgumentFlag = 4 // Indicates the logging level.
	FlagLogFacility    ArgumentFlag = 5 // Indicates the logging facility.
	FlagLogDestination ArgumentFlag = 6 // Indicates the logging target (stdout, file, syslog).

	// ToDo: Add as a common Flag in common/cli
	FlagForce  ArgumentFlag = 8 // Indicates force is to be used in all operations.
	FlagOutput ArgumentFlag = 9 // Identifies the target (path/file, url) where results should be written.

	// General attack Flags
	//		These Flags are common across many tools
	FlagConcurrency ArgumentFlag = 10 // Indicates attack concurrency (number of attackers)
	FlagDelay       ArgumentFlag = 11 // Indicates delay between attacks by any autonomous process.
	FlagDnsServer   ArgumentFlag = 12 // Specifies the DNS server to be used by attackers.
	FlagDomain      ArgumentFlag = 13 // Specifies a fully qualified domain name.
	FlagTimeout     ArgumentFlag = 14 // Provides a connection timeout for any attacker process.

	// Enumeration Flags
	FlagDepth ArgumentFlag = 40 // Indicates the depth of any enumeration attack.

	// Source-Control Flags
	FlagSource             ArgumentFlag = 60 // Declares a source mode (random, sequential, dictionary)
	FlagSourceDictionary   ArgumentFlag = 61 // Defines a dictionary path/filename (requires --source dictionary)
	FlagSourceMaxWordCount ArgumentFlag = 62 // Defines the maximum number of words to generate.
	FlagSourcePattern      ArgumentFlag = 63 // Specifies a regex patterns all generated words must match to be used.
	FlagSourceMinWordSize  ArgumentFlag = 64 // defines the minimum word size a generator should target.
	FlagSourceMaxWordSize  ArgumentFlag = 65 // defines the maximum word size a generator should target.

	// DNS tool-specific Flags
	FlagDnsRecordType ArgumentFlag = 900 //Dns record type identifier Flag.
)
