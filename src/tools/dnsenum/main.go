package main

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"os"
)

/*
	dnsEnum is a DNS enumerator project similar to the perl project of the same name.  This program aims to enumerate
	a DNS tree by brute force through either a dictionary, random or sequential string attack.
*/

func main() {
	//
	// Setup our program to run
	//
	var exit chan int = make(chan int, 1) // We will block until an exit code is written to this channel.
	config, exitProgram, err := ProcessSpecification(os.Args[1:])
	if err != nil {
		exit <- cli.ErrArgumentParseError
	}
	if exitProgram {
		exit <- cli.ErrSuccess
	}
	errors.Assert(config == nil, "Internal error nil config encountered.")

	var log logger.Logger // Global logger
	log.Setup(&config.Log)
	log.Debug(logger.EventInit)

	var feed source.Source	// Source Generator
	feed.Setup(&config.Source)		//		Pass the config object to the source.

	//var attack dnsEnumerator			  // Attacker Payload

	//var requestSent chan bool = make(chan bool, 1)

	/*
		var request dnsRequest.Request
		request.SetConcurrency(args.GetConcurrency())
		request.SetTimeout(args.GetTimeout())
		request.Domain(args.GetDomain())

		var collector response.Collector
		collector.Setup(ResponseBufferSz)
	*/
	/*
		go func() {
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
		}()
	*/
	/*go func() {
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
	}()*/

	<-exit
}
