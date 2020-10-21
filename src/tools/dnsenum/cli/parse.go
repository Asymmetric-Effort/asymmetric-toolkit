package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/types"
	"fmt"
	"regexp"
)

func (o *Configuration) Parse(cliArguments []string) bool {
	o.LoadDefault()
	expected := ExpectFlag
	lastFlag := NoFlag
	for _, args := range cliArguments {
		switch expected {
		case ExpectFlag:
			switch args {
			case "-h", "--help":
				showUsage()
				lastFlag = UsageFlag
				return ExitTerminate
			case "--version", "-v":
				showVersion()
				lastFlag = VersionFlag
				return ExitTerminate
			case "--concurrency":
				lastFlag = ConcurrencyFlag
				expected = ExpectValue
			case "--debug":
				lastFlag = DebugFlag
				expected = ExpectFlag
				o.Debug = true
			case "--delay":
				lastFlag = DelayFlag
				expected = ExpectValue
			case "--depth":
				lastFlag = DepthFlag
				expected = ExpectValue
			case "--dictionary":
				lastFlag = DictionaryFlag
				expected = ExpectValue
			case "--dnsServer":
				lastFlag = TargetServerFlag
				expected = ExpectValue
			case "--domain":
				lastFlag = DomainFlag
				expected = ExpectValue
			case "--force":
				lastFlag = ForceFlag
				expected = ExpectFlag
				o.Force = true
			case "--maxWordCount":
				lastFlag = MaxWordCountFlag
				expected = ExpectValue
			case "--mode":
				lastFlag = ModeFlag
				expected = ExpectValue
			case "--output":
				lastFlag = OutputFlag
				expected = ExpectValue
			case "--pattern":
				lastFlag = PatternFlag
				expected = ExpectValue
			case "--recordTypes":
				lastFlag = RecordTypesFlag
				expected = ExpectValue
			case "--timeout":
				lastFlag = TimeoutFlag
				expected = ExpectValue
			case "--wordSize":
				lastFlag = WordSizeFlag
				expected = ExpectValue
			default:
				errors.Fatal(1, fmt.Sprintf("Encountered unexpected argument: %s", args))
			}
		case ExpectValue:
			re:=regexp.MustCompile(`^--.+$`)
			if re.MatchString(args) {
				errors.Fatal(1, fmt.Sprintf("Expected value, not flag."))
			}
			switch lastFlag {
			case UsageFlag, VersionFlag:
				break
			case ConcurrencyFlag:
				o.Concurrency.Set(args)
				expected = ExpectFlag
			case DebugFlag:
				expected = ExpectFlag
				break
			case DepthFlag:
				expected = ExpectFlag
				o.Depth.Set(args)
			case DelayFlag:
				expected = ExpectFlag
				o.Delay.Set(args)
			case DictionaryFlag:
				expected = ExpectFlag
				o.Dictionary.Set(args)
			case DomainFlag:
				expected = ExpectFlag
				o.Domain.Set(args)
			case TargetServerFlag:
				expected = ExpectFlag
				o.TargetServer.Set(args)
			case ForceFlag:
				expected = ExpectFlag
				break
			case MaxWordCountFlag:
				o.MaxWordCount.Set(args)
				expected = ExpectFlag
			case ModeFlag:
				o.Mode.Set(args)
				expected = ExpectFlag
			case OutputFlag:
				o.Output.Set(args)
				expected = ExpectFlag
			case PatternFlag:
				o.Pattern.Set(args)
				expected = ExpectFlag
			case RecordTypesFlag:
				o.RecordTypes.Set(args)
				expected = ExpectFlag
			case TimeoutFlag:
				o.Timeout.Set(args)
				expected = ExpectFlag
			case WordSizeFlag:
				o.WordSize.Set(args)
				expected = ExpectFlag
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
			o.WordSize = DefaultWordSize
		}
	}
	if o.Output != "" {
		if !o.Force && o.Output.Exists() {
			fmt.Println("Output file exists.  Use --force to overwrite")
			return ExitTerminate
		}
	}
	if o.Depth > MaxDepth {
		fmt.Printf("Depth (--depth) exceeds maxDepth (%d)\n", MaxDepth)
		return ExitTerminate
	}
	return ExitParseOk
}
