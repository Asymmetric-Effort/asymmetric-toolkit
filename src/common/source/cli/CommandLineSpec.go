package sourcecli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/utils"
)

func CommandLineSpecification(config *cli.CommandLine) *cli.Specification {
	return &cli.Specification{

		FlagPrefix + FlagMaxWordCount: {
			NotRequired,
			ValueRequired,
			config.ValidateInteger(
				FlagMaxWordCount,
				FlagMaxWordCountDefault,
				FlagMaxWordCountLow,
				FlagMaxWordCountHigh),
			FlagMaxWordCountText,
		},

		FlagPrefix + FlagSource: {
			Required,
			ValueRequired,
			config.Enum(
				FlagSource,
				FlagSourceSequence,
				cli.EnumeratedValues{
					FlagSourceSequence,
					FlagSourceRandom,
					FlagSourceDictionary,
				}),
			FlagSourceText,
		},

		FlagPrefix + FlagDictionary: {
			NotRequired,
			ValueRequired,
			config.ValidateString(
				FlagDictionary,
				FlagDictionaryDefault,
				utils.RegExDotPlusMan),
			FlagDictionaryText,
		},

		FlagPrefix + FlagPattern: {
			NotRequired,
			ValueRequired,
			config.ValidateString(
				FlagPattern,
				FlagPatternDefault,
				utils.RegExDotPlusMan),
			FlagPatternText,
		},

		FlagPrefix + FlagWordsize: {
			NotRequired,
			ValueRequired,
			config.ValidateInteger(
				FlagWordsize,
				FlagWordSizeDefault,
				FlagWordSizeLow,
				FlagWordSizeHigh),
			FlagWordSizeText,
		},
	}
}
