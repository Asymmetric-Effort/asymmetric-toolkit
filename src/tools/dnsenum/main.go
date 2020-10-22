package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
)

func main() {
	var log logger.Logger
	var feed source.Source
	var exit chan bool = make(chan bool, 1)

	var config cli.Configuration //Load the common configuration (cli parser).
	config.Setup(&cli.Specification{
		cli.FlagShortHelp: {
			false,
			false,
			config.Bool(cli.FlagHelp, false),
			cli.FlagHelpText,
		},
		cli.FlagPrefix + cli.FlagHelp: { // Example: --help
			false,
			false,
			config.Bool(cli.FlagHelp, false),
			cli.FlagHelpText,
		},
		cli.FlagPrefix + cli.FlagVersion: {
			false,
			false,
			config.Bool(cli.FlagVersion, false),
			cli.FlagVersionText,
		},
		cli.FlagPrefix + cli.FlagDebug: {
			false,
			false,
			config.Bool(cli.FlagDebug, false),
			cli.FlagDebugText,
		},
		cli.FlagPrefix + cli.FlagForce: {
			false,
			false,
			config.Bool(cli.FlagForce, false),
			cli.FlagForceText,
		},
		cli.FlagPrefix + cli.FlagConcurrency: {
			false,
			true,
			config.Int(cli.FlagConcurrency, 10, 0, 100000),
			cli.FlagConcurrencyText,
		},
		cli.FlagPrefix + cli.FlagDelay: {
			false,
			true,
			config.Int(cli.FlagDelay, 0, 0, 100000),
			cli.FlagDelayText,
		},
		cli.FlagPrefix + cli.FlagDepth: {
			false,
			true,
			config.Int(cli.FlagDepth, 1, 1, 1000),
			cli.FlagDepthText,
		},
		cli.FlagPrefix + cli.FlagDictionary: {
			false,
			true,
			config.String(cli.FlagDictionary, "", ".+"),
			cli.FlagDictionaryText,
		},
		cli.FlagPrefix + cli.FlagTarget: {
			false,
			true,
			config.String(cli.FlagTarget, "", ".+"),
			cli.FlagTargetText,
		},
		cli.FlagPrefix + cli.FlagDomain: {
			false,
			true,
			config.String(cli.FlagDomain, "", ".+"),
			cli.FlagDomainText,
		},
		cli.FlagPrefix + cli.FlagMaxWordCount: {
			false,
			true,
			config.Int(cli.FlagMaxWordCount, 0, 1, 1000000000),
			cli.FlagMaxWordCountText,
		},
		cli.FlagPrefix + cli.FlagOutput: {
			false,
			true,
			config.String(cli.FlagOutput, "", ".+"),
			cli.FlagOutputText,
		},
		cli.FlagPrefix + cli.FlagPattern: {
			false,
			true,
			config.String(cli.FlagPattern, "", ".+"),
			cli.FlagPatternText,
		},
		cli.FlagPrefix + cli.FlagDNSRecordTypes: {
			false,
			true,
			config.String(cli.FlagDNSRecordTypes, "", ".+"),
			cli.FlagDNSRecordTypesText,
		},
		cli.FlagPrefix + cli.FlagTimeout: {
			false,
			true,
			config.Int(cli.FlagTimeout, 60, 1, 100000),
			cli.FlagTimeoutText,
		},
		cli.FlagPrefix + cli.FlagWordsize: {
			false,
			true,
			config.Int(cli.FlagWordsize, 0, 1, 1024),
			cli.FlagWordSizeText,
		},
		cli.FlagPrefix + cli.FlagSource: {
			false,
			true,
			config.Enum(cli.FlagSource, "sequence", "sequence", "random", "dictionary"),
			cli.FlagSourceText,
		},
	})

	log.Setup(&config)
	log.Debug("Main(): Logger is setup and buildConfig is loaded.")

	feed.Setup(&config, cli.SourceBufferSz, cli.DNSChars)

	//var requestSent chan bool = make(chan bool, 1)

	/*
		var request dnsRequest.Request
		request.SetConcurrency(args.GetConcurrency())
		request.SetTimeout(args.GetTimeout())
		request.Domain(args.GetDomain())

		var collector response.Collector
		collector.Setup(responseBufferSz)
	*/
	go func() {
		/*
				Consume the feed data to perform attacks.
			 *//*
			for source.HasData() {
				request.Query(source.Get())
				for _, recordType := range *args.GetTypes() {
					request.Type(&recordType)
					collector.Collect(request.Send())
					requestSent <- true
				}
			}
			requestSent <- false

		*/
	}()
	go func() {
		/*
			Consume server responses and write the report.
		*/
		/*
			var report Reporter.Report
			report.Setup(args.GetDomain(), args.GetDepth(), args.GetSource(),
				args.GetTypes(), args.GetConcurrency(), args.GetTimeout())
			for <-requestSent {
				report.Post(collector.Fetch())
			}
			report.Close()
			exit <- true
		*/
	}()
	<-exit
}
