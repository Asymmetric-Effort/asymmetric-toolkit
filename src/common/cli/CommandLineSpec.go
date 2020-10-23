package cli

/*
	This is the core set of command line arguments which should apply to all tools.
	Unless a tool overrides these terms, the specification will be consistent across
	all tools providing help, versioning and forced override flags.
*/
func CommandLineSpecification(config *CommandLine) *Specification {
	return &Specification{
		FlagShortHelp: {
			NotRequired,
			NoValueNeeded,
			config.ValidateBool(FlagHelp, false),
			FlagHelpText,
		},

		FlagPrefix + FlagHelp: { // Example: --help
			NotRequired,
			NoValueNeeded,
			config.ValidateBool(FlagHelp, false),
			FlagHelpText,
		},

		FlagPrefix + FlagVersion: {
			NotRequired,
			NoValueNeeded,
			config.ValidateBool(FlagVersion, false),
			FlagVersionText,
		},

		FlagPrefix + FlagForce: {
			NotRequired,
			NoValueNeeded,
			config.ValidateBool(FlagForce, false),
			FlagForceText,
		},
	}
}
