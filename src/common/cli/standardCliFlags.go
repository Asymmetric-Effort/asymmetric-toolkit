package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/utils"
	logLevel "github.com/combat-diver-foundation/common/logger/log_level"
)

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

	FlagDictionary     = "dictionary"
	FlagDictionaryText = "Specify the path/filename of a given dictionary file.  This dictionary file will be " +
		"in the proprietary file format implemented by the toolkit."
	FlagDictionaryDefault = ""

	FlagDomain        = "domain"
	FlagDomainText    = "Specify a fully-qualified or partial domain name string."
	FlagDomainDefault = ""

	FlagForce     = "force"
	FlagForceText = "Specify that any operation which would otherwise fail should proceed with force.  " +
		"This may mean files are overwritten or other actions are taken without warning.  Use caution " +
		"and read the documentation for each tool specifically."

	FlagHelp     = "help"
	FlagHelpText = "Show this help documentation."

	FlagLogDestination        = "logdestination"
	FlagLogDestinationText    = "Specify the destination for log activity (e.g. stdout, file, syslog)"
	FlagLogDestinationDefault = "stdout"

	FlagLogFile        = "logfile"
	FlagLogFileText    = "Specify the path/filename of a log file where file is the destination."
	FlagLogFileDefault = "asymtoolkit.log"

	FlagLogLevel        = "loglevel"
	FlagLogLevelText    = "Specify the logging level (string)."
	FlagLogLevelDefault = logLevel.InfoText

	FlagLogServer        = "logserver"
	FlagLogServerText    = "Specify the FQDN and port for sending logs to a syslog server."
	FlagLogServerDefault = "udp:127.0.0.1:514"
	//ToDo: add TLS options for syslog.

	FlagMaxWordCount     = "maxwordcount"
	FlagMaxWordCountText = "Specify the maximum number of words to generate from the standard source generator " +
		"using either the sequential, random or dictionary generator."
	FlagMaxWordCountDefault = 0
	FlagMaxWordCountLow     = 1
	FlagMaxWordCountHigh    = 1000000000

	FlagSource     = "source"
	FlagSourceText = "Specify the type of word source to be used for an attack (sequence, random or dictionary) " +
		"where each source will generate a series of 'word' strings used as specified in the tool documentation."

	FlagOutput        = "output"
	FlagOutputText    = "Specify the output path/filename for the report to be generated following tool termination."
	FlagOutputDefault = ""

	FlagPattern        = "pattern"
	FlagPatternText    = "Specify a regular expression pattern as a string."
	FlagPatternDefault = utils.RegExDotPlusMan

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

	FlagWordsize     = "wordsize"
	FlagWordSizeText = "Specify the number of characters in a word generated by the Sequence or Random " +
		"source generators."
	FlagWordSizeDefault = 0
	FlagWordSizeLow     = 1
	FlagWordSizeHigh    = 1024
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
