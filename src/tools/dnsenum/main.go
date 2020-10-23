package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"asymmetric-effort/asymmetric-toolkit/src/common/utils"
	"fmt"
)

func main() {
	var err error
	var log logger.Logger
	var feed source.Source
	var exit chan bool = make(chan bool, 1)

	var config cli.Configuration //Load the common configuration (cli parser).
	err = config.Setup(&cli.Specification{
		//ToDo: Move to the attacker package
		cli.FlagPrefix + cli.FlagConcurrency: {
			cli.NotRequired,
			cli.ValueRequired,
			config.Int(
				cli.FlagConcurrency,
				cli.FlagConcurrencyDefault,
				cli.FlagConcurrencyLow,
				cli.FlagConcurrencyHigh),
			cli.FlagConcurrencyText,
		},

		//ToDo: Move to the attacker package
		cli.FlagPrefix + cli.FlagDelay: {
			cli.NotRequired,
			cli.ValueRequired,
			config.Int(
				cli.FlagDelay,
				cli.FlagDelayDefault,
				cli.FlagDelayLow,
				cli.FlagDelayHigh),
			cli.FlagDelayText,
		},

		//ToDo: Move to the attacker package
		cli.FlagPrefix + cli.FlagDepth: {
			cli.NotRequired,
			cli.ValueRequired,
			config.Int(
				cli.FlagDepth,
				cli.FlagDepthDefault,
				cli.FlagDepthLow,
				cli.FlagDepthHigh),
			cli.FlagDepthText,
		},

		//ToDo: Move to the attacker package
		cli.FlagPrefix + cli.FlagTarget: {
			cli.NotRequired,
			cli.ValueRequired,
			config.String(
				cli.FlagTarget,
				cli.FlagTargetDefaultDns,
				utils.RegExDotPlusMan),
			cli.FlagTargetText,
		},

		//ToDo: Move to the attacker package
		cli.FlagPrefix + cli.FlagDomain: {
			cli.NotRequired,
			cli.ValueRequired,
			config.String(
				cli.FlagDomain,
				cli.FlagDomainDefault,
				utils.RegExDotPlusMan),
			cli.FlagDomainText,
		},

		//ToDo: Move to the reporter package
		cli.FlagPrefix + cli.FlagOutput: {
			cli.NotRequired,
			cli.ValueRequired,
			config.String(
				cli.FlagOutput,
				cli.FlagOutputDefault,
				utils.RegExDotPlusMan),
			cli.FlagOutputText,
		},


		cli.FlagPrefix + cli.FlagDNSRecordTypes: {
			cli.NotRequired,
			cli.ValueRequired,
			config.String(
				cli.FlagDNSRecordTypes,
				cli.DefaultDNSRecordTypes,
				utils.RegExDotPlusMan),
			cli.FlagDNSRecordTypesText,
		},

		//ToDo: Move to the attacker package
		cli.FlagPrefix + cli.FlagTimeout: {
			cli.NotRequired,
			cli.ValueRequired,
			config.Int(
				cli.FlagTimeout,
				cli.FlagTimeoutDefault,
				cli.FlagTimeoutLow,
				cli.FlagTimeoutHigh),
			cli.FlagTimeoutText,
		},
	})
	errors.Assert(err == nil, fmt.Sprintf("%v", err))

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
