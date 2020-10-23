package cli

/*
	This is the core set of command line arguments which should apply to all tools.
	Unless a tool overrides these terms, the specification will be consistent across
	all tools providing help, versioning and forced override flags.
*/
func CommandLineSpecification(config *Configuration) *Specification {
	return &Specification{
		FlagShortHelp: {
			NotRequired,
			NoValueNeeded,
			config.Bool(FlagHelp, false),
			FlagHelpText,
		},

		FlagPrefix + FlagHelp: { // Example: --help
			NotRequired,
			NoValueNeeded,
			config.Bool(FlagHelp, false),
			FlagHelpText,
		},

		FlagPrefix + FlagVersion: {
			NotRequired,
			NoValueNeeded,
			config.Bool(FlagVersion, false),
			FlagVersionText,
		},

		FlagPrefix + FlagForce: {
			NotRequired,
			NoValueNeeded,
			config.Bool(FlagForce, false),
			FlagForceText,
		},
	}
}
