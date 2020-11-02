package cli

type ArgumentFlag int

const ( //Reserved 0...1000 for core use.  1000+ are user-defined.
	noFlag      ArgumentFlag = 0
	flagHelp    ArgumentFlag = 1
	flagVersion ArgumentFlag = 2
)