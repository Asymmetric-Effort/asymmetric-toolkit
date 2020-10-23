package cli

/*
	Standardized CLI flags.

		Intention: We should reuse flags for all tools in the toolkit and standardize their expected
				   behavior.  This block of constants aims to do just that.
*/

const (
	FlagConcurrency        = "concurrency"
	FlagConcurrencyText    = "Specify the number of concurrent operations to perform in the attack."
	FlagConcurrencyDefault = 10
	FlagConcurrencyLow     = 1
	FlagConcurrencyHigh    = 100000

	FlagDebug     = "debug"
	FlagDebugText = "Enable debug (verbose) logging in the logger facilities."

	FlagDelay        = "delay"
	FlagDelayText    = "Specify the number of seconds to wait between attack operations."
	FlagDelayDefault = 0
	FlagDelayLow     = 0
	FlagDelayHigh    = 100000

	FlagDepth     = "depth"
	FlagDepthText = "Specify the depth of the attack (e.g. number of iterations or layers into a target) " +
		"an attack will go until the tool stops.  For example, in the case of dnsenum this flag will indicate " +
		"the number of subdomain levels which will be enumerated by brute force."
	FlagDepthDefault = 1
	FlagDepthLow     = 1
	FlagDepthHigh    = 1000

	FlagDomain        = "domain"
	FlagDomainText    = "Specify a fully-qualified or partial domain name string."
	FlagDomainDefault = ""

	FlagForce     = "force"
	FlagForceText = "Specify that any operation which would otherwise fail should proceed with force.  " +
		"This may mean files are overwritten or other actions are taken without warning.  Use caution " +
		"and read the documentation for each tool specifically."

	FlagHelp     = "help"
	FlagHelpText = "Show this help documentation."

	FlagOutput        = "output"
	FlagOutputText    = "Specify the output path/filename for the report to be generated following tool termination."
	FlagOutputDefault = ""

	FlagDNSRecordTypes     = "dnsrecordtypes"
	FlagDNSRecordTypesText = "DNS Record types list."

	FlagTarget           = "target"
	FlagTargetText       = "Specify an attack target (string)"
	FlagTargetDefaultDns = "udp:127.0.0.1:53"

	FlagTimeout        = "timeout"
	FlagTimeoutText    = "Specify a connection timeout in (integer) seconds."
	FlagTimeoutDefault = 60
	FlagTimeoutLow     = 1
	FlagTimeoutHigh    = 100000

	FlagVersion     = "version"
	FlagVersionText = "Show the toolkit/application version.  Note: All tools from a single build will " +
		"have the same semantic version number (0.0.0)."
)
const (
	/*
		miscellaneous parts.
	*/
	FlagPrefix    = "--"
	FlagShortHelp = "-h"

	//Require argument
	Required    = true
	NotRequired = false

	//ExpectValue argument
	ValueRequired = true
	NoValueNeeded = false
)
