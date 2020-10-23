package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"fmt"
	"regexp"
)

func (o *Configuration) Parse(cliArguments []string) bool {
	o.LoadDefault()
	expected := expectFlag
	lastFlag := noFlag
	for _, args := range cliArguments {
		switch expected {
		case expectFlag:
			switch args {
			case "-h", "--help":
				showUsage()
				lastFlag = usageFlag
				return ExitTerminate
			case "--version", "-v":
				showVersion()
				lastFlag = versionFlag
				return ExitTerminate
			case "--concurrency":
				lastFlag = concurrencyFlag
				expected = expectValue
			case "--debug":
				lastFlag = debugFlag
				expected = expectFlag
				o.Debug = true
			case "--delay":
				lastFlag = delayFlag
				expected = expectValue
			case "--depth":
				lastFlag = depthFlag
				expected = expectValue
			case "--dictionary":
				lastFlag = dictionaryFlag
				expected = expectValue
			case "--dnsServer":
				lastFlag = targetServerFlag
				expected = expectValue
			case "--domain":
				lastFlag = domainFlag
				expected = expectValue
			case "--force":
				lastFlag = forceFlag
				expected = expectFlag
				o.Force = true
			case "--maxWordCount":
				lastFlag = maxWordCountFlag
				expected = expectValue
			case "--mode":
				lastFlag = modeFlag
				expected = expectValue
			case "--output":
				lastFlag = outputFlag
				expected = expectValue
			case "--pattern":
				lastFlag = patternFlag
				expected = expectValue
			case "--recordTypes":
				lastFlag = recordTypesFlag
				expected = expectValue
			case "--timeout":
				lastFlag = timeoutFlag
				expected = expectValue
			case "--wordSize":
				lastFlag = wordSizeFlag
				expected = expectValue
			default:
				errors.Fatal(1, fmt.Sprintf("Encountered unexpected argument: %s", args))
			}
		case expectValue:
			re:=regexp.MustCompile(`^--.+$`)
			if re.MatchString(args) {
				errors.Fatal(1, fmt.Sprintf("Expected value, not flag."))
			}
			switch lastFlag {
			case usageFlag, versionFlag:
				break
			case concurrencyFlag:
				o.Concurrency.Set(args)
				expected = expectFlag
			case debugFlag:
				expected = expectFlag
				break
			case depthFlag:
				expected = expectFlag
				o.Depth.Set(args)
			case delayFlag:
				expected = expectFlag
				o.Delay.Set(args)
			case dictionaryFlag:
				expected = expectFlag
				o.Dictionary.Set(args)
			case domainFlag:
				expected = expectFlag
				o.Domain.Set(args)
			case targetServerFlag:
				expected = expectFlag
				o.TargetServer.Set(args)
			case forceFlag:
				expected = expectFlag
				break
			case maxWordCountFlag:
				o.MaxWordCount.Set(args)
				expected = expectFlag
			case modeFlag:
				o.Mode.Set(args)
				expected = expectFlag
			case outputFlag:
				o.Output.Set(args)
				expected = expectFlag
			case patternFlag:
				o.Pattern.Set(args)
				expected = expectFlag
			case recordTypesFlag:
				o.RecordTypes.Set(args)
				expected = expectFlag
			case timeoutFlag:
				o.Timeout.Set(args)
				expected = expectFlag
			case wordSizeFlag:
				o.WordSize.Set(args)
				expected = expectFlag
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
	if o.Domain.Get() == "" {
		fmt.Println("Missing domain (required).  Use --domain <string> to specify.")
		return ExitTerminate
	}

	if o.Mode.Get() == types.NotSet {
		fmt.Println("Missing mode (required).  Use --mode <sequence|random|dictionary> to specify.")
		return ExitTerminate
	}

	if o.TargetServer.Get() == "" {
		fmt.Println("Missing dnsServer (required).  Use --dnsServer <udp|tcp>:<ipaddr>:<port> to specify.")
		return ExitTerminate
	}

	if o.Mode.IsDictionary() {
		if o.Dictionary == "" {
			fmt.Println("Missing Dictionary.  Use --dictionary <string> when --mode dictionary is used.")
			return ExitTerminate
		}
		if !o.Dictionary.Exists() {
			fmt.Printf("Dictionary file not found (%s)\n", o.Dictionary)
			return ExitTerminate
		}
		if o.WordSize >= 0 {
			fmt.Printf("In Dictionary mode --wordSize is not allowed")
			return ExitTerminate
		}
		if o.MaxWordCount >= 0 {
			fmt.Printf("In Dictionary mode --maxWordCount is not allowed")
			return ExitTerminate
		}
	}
	if o.Mode.IsRandom() || o.Mode.IsSequence() {
		if o.Dictionary != "" {
			fmt.Println("Do not use --dictionary <file> with random or sequential mode.")
			return ExitTerminate
		}
		if o.MaxWordCount == 0 {
			o.MaxWordCount = types.PositiveInteger(int(o.WordSize) * len(DnsChars))
		}
		if o.WordSize == 0 {
			o.WordSize = defaultWordSize
		}
	}
	if o.Output != "" {
		if !o.Force && o.Output.Exists() {
			fmt.Println("Output file exists.  Use --force to overwrite")
			return ExitTerminate
		}
	}
	if o.Depth > maxDepth {
		fmt.Printf("Depth (--depth) exceeds maxDepth (%d)\n", maxDepth)
		return ExitTerminate
	}
	return ExitParseOk
}
