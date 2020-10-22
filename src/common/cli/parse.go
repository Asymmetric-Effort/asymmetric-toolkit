package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/cli"
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/sourcetype"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"fmt"
	"regexp"
)

func (o *cli.Configuration) Parse(cliArguments []string) bool {
	o.LoadDefault()
	expected := cli.ExpectFlag
	lastFlag := cli.NoFlag
	for _, args := range cliArguments {
		switch expected {
		case cli.ExpectFlag:
			switch args {
			case "-h", "--help":
				ShowUsage()
				lastFlag = cli.UsageFlag
			case "--version", "-v":
				ShowVersion()
				lastFlag = cli.VersionFlag
			case "--concurrency":
				lastFlag = cli.ConcurrencyFlag
				expected = cli.ExpectValue
			case "--debug":
				lastFlag = cli.DebugFlag
				expected = cli.ExpectFlag
				o.Debug = true
			case "--delay":
				lastFlag = cli.DelayFlag
				expected = cli.ExpectValue
			case "--depth":
				lastFlag = cli.DepthFlag
				expected = cli.ExpectValue
			case "--dictionary":
				lastFlag = cli.DictionaryFlag
				expected = cli.ExpectValue
			case "--dnsServer":
				lastFlag = cli.TargetServerFlag
				expected = cli.ExpectValue
			case "--domain":
				lastFlag = cli.DomainFlag
				expected = cli.ExpectValue
			case "--force":
				lastFlag = cli.ForceFlag
				expected = cli.ExpectFlag
				o.Force = true
			case "--maxWordCount":
				lastFlag = cli.MaxWordCountFlag
				expected = cli.ExpectValue
			case "--mode":
				lastFlag = cli.ModeFlag
				expected = cli.ExpectValue
			case "--output":
				lastFlag = cli.OutputFlag
				expected = cli.ExpectValue
			case "--pattern":
				lastFlag = cli.PatternFlag
				expected = cli.ExpectValue
			case "--recordTypes":
				lastFlag = cli.RecordTypesFlag
				expected = cli.ExpectValue
			case "--timeout":
				lastFlag = cli.TimeoutFlag
				expected = cli.ExpectValue
			case "--wordSize":
				lastFlag = cli.WordSizeFlag
				expected = cli.ExpectValue
			default:
				errors.Fatal(1, fmt.Sprintf("Encountered unexpected argument: %s", args))
			}
		case cli.ExpectValue:
			func() {
				re := regexp.MustCompile(`^--.+$`)
				if re.MatchString(args) {
					errors.Fatal(1, "Expected value, not flag.")
				}
			}()

			switch lastFlag {
			case cli.NoFlag, cli.UsageFlag:
				ShowUsage()
			case cli.VersionFlag:
				break
			case cli.ConcurrencyFlag:
				o.Concurrency.Set(args)
				expected = cli.ExpectFlag
			case cli.DebugFlag:
				expected = cli.ExpectFlag
			case cli.DepthFlag:
				expected = cli.ExpectFlag
				o.Depth.Set(args)
			case cli.DelayFlag:
				expected = cli.ExpectFlag
				o.Delay.Set(args)
			case cli.DictionaryFlag:
				expected = cli.ExpectFlag
				o.Dictionary.Set(args)
			case cli.DomainFlag:
				expected = cli.ExpectFlag
				o.Domain.Set(args)
			case cli.TargetServerFlag:
				expected = cli.ExpectFlag
				o.TargetServer.Set(args)
			case cli.ForceFlag:
				expected = cli.ExpectFlag
			case cli.MaxWordCountFlag:
				o.MaxWordCount.Set(args)
				expected = cli.ExpectFlag
			case cli.ModeFlag:
				o.Mode.Set(args)
				expected = cli.ExpectFlag
			case cli.OutputFlag:
				o.Output.Set(args)
				expected = cli.ExpectFlag
			case cli.PatternFlag:
				o.Pattern.Set(args)
				expected = cli.ExpectFlag
			case cli.RecordTypesFlag:
				o.RecordTypes.Set(args)
				expected = cli.ExpectFlag
			case cli.TimeoutFlag:
				o.Timeout.Set(args)
				expected = cli.ExpectFlag
			case cli.WordSizeFlag:
				o.WordSize.Set(args)
				expected = cli.ExpectFlag
			default:
				panic("invalid flag")
			}
		default:
			panic(fmt.Sprintf("Expected either expectFlag or expectValue.  Encountered: %v", expected))
		}
	} /* end for */
	//
	//Perform a final validation...
	//
	if lastFlag == cli.UsageFlag {
		ShowUsage()
		return cli.ExitTerminate
	}
	if lastFlag == cli.VersionFlag {
		ShowVersion()
		return cli.ExitTerminate
	}
	if o.Domain.Get() == "" {
		fmt.Println("Missing domain (required).  Use --domain <string> to specify.")
		return cli.ExitTerminate
	}

	if o.Mode.Get() == sourcetype.NotSet {
		fmt.Println("Missing mode (required).  Use --mode <sequence|random|dictionary> to specify.")
		return cli.ExitTerminate
	}

	if o.TargetServer.Get() == "" {
		fmt.Println("Missing dnsServer (required).  Use --dnsServer <udp|tcp>:<ipaddr>:<port> to specify.")
		return cli.ExitTerminate
	}

	if o.Mode.IsDictionary() {
		if o.Dictionary == "" {
			fmt.Println("Missing Dictionary.  Use --dictionary <string> when --mode dictionary is used.")
			return cli.ExitTerminate
		}
		if !o.Dictionary.Exists() {
			fmt.Printf("Dictionary file not found (%s)\n", o.Dictionary)
			return cli.ExitTerminate
		}
		if o.WordSize >= 0 {
			fmt.Printf("In Dictionary mode --wordSize is not allowed")
			return cli.ExitTerminate
		}
		if o.MaxWordCount >= 0 {
			fmt.Printf("In Dictionary mode --maxWordCount is not allowed")
			return cli.ExitTerminate
		}
	}
	if o.Mode.IsRandom() || o.Mode.IsSequence() {
		if o.Dictionary != "" {
			fmt.Println("Do not use --dictionary <file> with random or sequential mode.")
			return cli.ExitTerminate
		}
		if o.MaxWordCount == 0 {
			o.MaxWordCount = types.PositiveInteger(int(o.WordSize) * len(cli.DNSChars))
		}
		if o.WordSize == 0 {
			o.WordSize = cli.DefaultWordSize
		}
	}
	if o.Output != "" {
		if !o.Force && o.Output.Exists() {
			fmt.Println("Output file exists.  Use --force to overwrite")
			return cli.ExitTerminate
		}
	}
	if o.Depth > cli.MaxDepth {
		fmt.Printf("Depth (--depth) exceeds maxDepth (%d)\n", cli.MaxDepth)
		return cli.ExitTerminate
	}
	return cli.ExitParseOk
}
