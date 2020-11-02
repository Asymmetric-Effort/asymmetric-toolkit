package cli

type ArgumentType int

const (
	None    ArgumentType = 0
	String  ArgumentType = 1
	Integer ArgumentType = 2
	Float   ArgumentType = 3
	Boolean ArgumentType = 4
	Enum    ArgumentType = 5
)

/*
"version":cli.ArgumentDescriptor{
				cli.None,
				"",
				"Display version",
				func() {},
			},
			"debug":{

			},
 */