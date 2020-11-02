package cli

const (
	expectFlag  int = 0
	expectValue int = 1
	expectEnd   int = 2
)
const (
	noFlag   int = 0
	flagHelp int = 1
	flagVersion int = 2
)

const (
	helpArgLong  = "--help"
	helpArgShort = "-h"

	versionArgLong = "--version"
	versionArgShort = "-v"
)

const (
	ErrSuccess = 0
	ErrArgumentParseError=1
)