package cli

type ArgFlags int

const (
	noFlag            ArgFlags = 0
	domainFlag        ArgFlags = 1
	modeFlag          ArgFlags = 2
	depthFlag         ArgFlags = 3
	patternFlag       ArgFlags = 4
	outputFlag        ArgFlags = 5
	dictionaryFlag    ArgFlags = 6
	concurrencyFlag   ArgFlags = 7
	timeoutFlag       ArgFlags = 8
	delayFlag         ArgFlags = 9
	targetServersFlag ArgFlags = 10
	recordTypesFlag   ArgFlags = 11
	debugFlag         ArgFlags = 12
	forceFlag         ArgFlags = 13
	usageFlag         ArgFlags = 14
	versionFlag       ArgFlags = 15
	wordSizeFlag      ArgFlags = 16
	maxWordCountFlag  ArgFlags = 17
)
