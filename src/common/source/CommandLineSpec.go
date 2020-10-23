package source

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/utils"
)

func CommandLineSpecification(config *cli.Configuration) *cli.Specification {
	return &cli.Specification{

		cli.FlagPrefix + FlagMaxWordCount: {
			cli.NotRequired,
			cli.ValueRequired,
			config.Int(
				FlagMaxWordCount,
				FlagMaxWordCountDefault,
				FlagMaxWordCountLow,
				FlagMaxWordCountHigh),
			FlagMaxWordCountText,
		},

		cli.FlagPrefix + FlagSource: {
			cli.Required,
			cli.ValueRequired,
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

		cli.FlagPrefix + FlagDictionary: {
			cli.NotRequired,
			cli.ValueRequired,
			config.String(
				FlagDictionary,
				FlagDictionaryDefault,
				utils.RegExDotPlusMan),
			FlagDictionaryText,
		},

		cli.FlagPrefix + FlagPattern: {
			cli.NotRequired,
			cli.ValueRequired,
			config.String(
				FlagPattern,
				FlagPatternDefault,
				utils.RegExDotPlusMan),
			FlagPatternText,
		},

		cli.FlagPrefix + FlagWordsize: {
			cli.NotRequired,
			cli.ValueRequired,
			config.Int(
				FlagWordsize,
				FlagWordSizeDefault,
				FlagWordSizeLow,
				FlagWordSizeHigh),
			FlagWordSizeText,
		},
	}
}
