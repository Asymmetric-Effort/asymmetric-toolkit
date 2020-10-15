package cli

const (
	Usage            string = `

	Asymmetric-Toolkit (dnsEnum): dnsEnumeration tool.
	(c) 2020 Sam Caldwell.  All Rights Reserved.

	Usage:

		Get Help:
			dnsEnum -h | --help

		Get Current Version:
			dnsEnum --version

		Brute Force with Random or Sequential Inputs:
			dnsEnum --domain <root_domain> --mode <random|sequence> [--depth <int>] \
					[--pattern /regex/] [--output <file>] [--concurrency <int>]  \
					[--timeout <int>]  [--delay <int>]  [--TargetServers <string>] \
					[--recordTypes <str>] [--debug] --MaxWordCount <int> --WordSize <int>

		Brute Force with Dictionary File:
			dnsEnum --domain <root_domain> --mode <dictionary> --dictionary <file> \
					[--depth <int>] [--pattern /regex/] [--output <file>] [--concurrency <int>] \
					[--timeout <int>]  [--delay <int>]  [--TargetServers <string>] \
					[--recordTypes <str>] [--debug]


	Arguments:

		--concurrency <int>	(default: 1)
			Number of concurrent queries to send to the target servers in round-robin fashion.

		--debug
			Print verbose logs/output.

		--delay <int>	(default: 0)
			Number of seconds to wait between queries.

		--depth <int> (default: 1)
			This is the number of levels deep the dns should be scanned. For each level of subdomain
			where a positive result is found, a successive scan of these positives will be run until
			the scanner has probed this number of levels deep or until no new positives are found in
			a given pass.

		--dictionary <path/filename> (Default: "")
			This is the path/filename to the dictionary file which will be used.  The dictionary file
			must follow the Asymmetric Toolkit dictionary file standard.

		--TargetServers <string> (Default: Uses Local resolver)
			A comma-delimited list of DNS servers to target in round robin fashion so as to avoid rate limiting
			or to evade detection.

		--domain <domainString> (Required)
			The root domain (e.g. my.domain.tld) to which we will prepend subdomains as part of our scan.

		--MaxWordCount <int> (Optional for Sequence, Random; not allowed for Dictionary)
			Maximum number of words to be generated.  Default is the number of allowed characters times
			the --WordSize value.

		--mode <mode> (Required)
			The source feed generator mode (random, sequence, dictionary).  This determines how we will
			generate strings to be prepended in our brute force attack.

		--output <path/filename> (Default: stdout)
			This is the path/filename where results should be written.  If not provided, the results
			will go to stdout.

		--pattern <regex> (Default: .+)
			This regular expression will be applied to all generated inputs as a filter, allowing the
			program to test only those generated inputs which match the pattern.

		--recordTypes <list> (Default: "A,AAAA,MX,CNAME,NS,TXT,SOA,SRV")
			comma-delimited list of strings identifying DNS record types (e.g. A, AAAA, CNAME, etc).

		--timeout <int> (Default: 60s)
			Number of seconds until query timeout.

		--WordSize <int> (Required for Sequence, Random; not allowed for Dictionary) 
		  	Maximum length of words to be generated.

`
)