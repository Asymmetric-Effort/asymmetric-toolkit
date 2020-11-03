package cli

/*
	Common flag identifiers intended to promote standardization
	across the asymmetric toolkit.  For all reusable commandline
	flags, we reserve 0-1000.  For program-specific flags not
	shared across the tool suite, we reserve numbers > 1000.
*/
type ArgumentFlag int

const ( //Reserved 0...1000 for core use.  1000+ are program-defined.
	noFlag      ArgumentFlag = 0
	flagHelp    ArgumentFlag = 1
	flagVersion ArgumentFlag = 2
)
