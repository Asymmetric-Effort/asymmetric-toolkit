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
	noFlag      ArgumentFlag = 0 // noop
	flagHelp    ArgumentFlag = 1 // Displays usage
	flagVersion ArgumentFlag = 2 // Displays version string
	flagDebug   ArgumentFlag = 3 // Indicates debug logging.

	// ToDo: Add as a common flag in common/cli
	flagForce  ArgumentFlag = 4 // Indicates force is to be used in all operations.
	flagOutput ArgumentFlag = 5 // Identifies the target (path/file, url) where results should be written.

	// General attack flags
	//		These flags are common across many tools
	flagConcurrency ArgumentFlag = 10 // Indicates attack concurrency (number of attackers)
	flagDelay       ArgumentFlag = 11 // Indicates delay between attacks by any autonomous process.
	flagDnsServer   ArgumentFlag = 12 // Specifies the DNS server to be used by attackers.
	flagDomain      ArgumentFlag = 13 // Specifies a fully qualified domain name.
	flagTimeout     ArgumentFlag = 14 // Provides a connection timeout for any attacker process.

	// Enumeration flags
	flagDepth ArgumentFlag = 40 // Indicates the depth of any enumeration attack.

	// Source-Control flags
	flagSource             ArgumentFlag = 60 // Declares a source mode (random, sequential, dictionary)
	flagSourceDictionary   ArgumentFlag = 61 // Defines a dictionary path/filename (requires --source dictionary)
	flagSourceMaxWordCount ArgumentFlag = 62 // Defines the maximum number of words to generate.
	flagSourcePattern      ArgumentFlag = 63 // Specifies a regex patterns all generated words must match to be used.
	flagSourceMinWordSize  ArgumentFlag = 64 // defines the minimum word size a generator should target.
	flagSourceMaxWordSize  ArgumentFlag = 65 // defines the maximum word size a generator should target.

	// DNS tool-specific flags
	flagDnsRecordType ArgumentFlag = 900 //Dns record type identifier flag.
)
