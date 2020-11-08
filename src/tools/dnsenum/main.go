package main

/*
	dnsEnum is a DNS enumerator project similar to the perl project of the same name.  This program aims to enumerate
	a DNS tree by brute force through either a dictionary, random or sequential string attack.
*/

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/logger"
	"asymmetric-effort/asymmetric-toolkit/src/common/source"
	"asymmetric-effort/asymmetric-toolkit/src/tools/dnsenum/deprecated_cli"
	"fmt"
	"os"
)

const (
	dnsRecordTypes = "A,AAAA,MX,TXT,SOA,NS,CNAME"
)

func main() {
	//
	// Setup our program to run
	//
	var exit chan int = make(chan int, 1) // We will block until an exit code is written to this channel.
	var config *cli.ArgumentMap			  // Runtime program configuration.
	var log logger.Logger				  // Global logger
	var feed source.Source				  // Source Generator
	//var attack dnsEnumerator			  // Attacker Payload
	func() {
		var spec cli.Specification = cli.Specification{
			Author:      "Sam Caldwell",
			AuthorEmail: "mail@samcaldwell.net",
			ProgramName: "Asymmetric Effort",
			Copyright:   "(c) 2018 Sam Caldwell.  All Rights Reserved.",
			Version:     "0.0.1",
			Argument: map[string]cli.ArgumentDescriptor{
				"recordTypes": {
					1000,
					cli.String,
					dnsRecordTypes,
					"Specify the DNS record types to be enumerated.",
					cli.ParseList(","),
					cli.ExpectValue,
				},
			},
		}
		var ui cli.CommandLine
		exitProgram, err := ui.Parse(&spec, os.Args[1:])
		if err != nil {
			fmt.Printf("Error:%v", err)
			exit <- cli.ErrArgumentParseError
		} else {
			if exitProgram {
				fmt.Printf("%v", err)
				exit <- cli.ErrSuccess
			} else {
				//
				// ToDo: get properties out of ui and free that memory.
				//
				config=&ui.Arguments
			}
		}
	}()
	log.Setup(&ui)
	log.Debug("Main(): Logger is setup and config is loaded.")

	feed.Setup(&ui, deprecated_cli.SourceBufferSz, deprecated_cli.DnsChars)

	//var requestSent chan bool = make(chan bool, 1)

	/*
		var request dnsRequest.Request
		request.SetConcurrency(args.GetConcurrency())
		request.SetTimeout(args.GetTimeout())
		request.Domain(args.GetDomain())

		var collector response.Collector
		collector.Setup(responseBufferSz)
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
