package cli

type ArgFlags int

const (
	NoFlag           ArgFlags = 0
	DomainFlag       ArgFlags = 1
	ModeFlag         ArgFlags = 2
	DepthFlag        ArgFlags = 3
	PatternFlag      ArgFlags = 4
	OutputFlag       ArgFlags = 5
	DictionaryFlag   ArgFlags = 6
	ConcurrencyFlag  ArgFlags = 7
	TimeoutFlag      ArgFlags = 8
	DelayFlag        ArgFlags = 9
	TargetServerFlag ArgFlags = 10
	RecordTypesFlag  ArgFlags = 11
	DebugFlag        ArgFlags = 12
	ForceFlag        ArgFlags = 13
	UsageFlag        ArgFlags = 14
	VersionFlag      ArgFlags = 15
	WordSizeFlag     ArgFlags = 16
	MaxWordCountFlag ArgFlags = 17
)
