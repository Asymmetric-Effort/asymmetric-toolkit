package cli

type ArgumentDescriptor struct {
	Type    ArgumentType // An expected data type.
	Default string       // The default value
	Help    string       // Help text
	Parse   ParserFunction // Argument-specific parser which will set the internal string state or return an error.
}

type ParserFunction func() error

func NoopParser()error {return nil}