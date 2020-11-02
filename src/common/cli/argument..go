package cli

type Argument struct {
	Type  ArgumentType
	Value string		// This is the parsed, validated but not decoded (i.e. string) value parsed from the cli.
}
