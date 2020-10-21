package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/cli"
	"os"
)

func main() {
	/*
		Parse/validate all commandline arguments and expose them as a single config
		object
	*/
	var config cli.Configuration
	var log logger.Logger
	var feed source.Source
	var exit chan bool = make(chan bool, 1)

	if config.Parse(os.Args) {
		os.Exit(1)
	}
	log.Setup(&config)
	log.Debugf("Main(): Logger is setup and config is loaded.")

	feed.Setup(&config,cli.SourceBufferSz,cli.DnsChars)

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
