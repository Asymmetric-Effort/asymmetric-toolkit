package cli

/*
	Common constants for the CommandLine processor.
*/
const (
	ErrSuccess            = 0
	ErrArgumentParseError = 1

	ErrBadFilename    = 10 // Bad or empty filename
	ErrFileNotFound   = 11 // Expected file is not found.
	ErrFileOpenFailed = 12 // Failed to open file
	ErrFileRead       = 13 // File read operation error.

	MissingArgumentsMessage = "Missing arguments (use --help for usage)"
	)
