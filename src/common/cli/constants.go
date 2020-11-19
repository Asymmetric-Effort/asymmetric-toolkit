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

	ErrMsgMissingArguments = "\n\tMissing arguments (use --help for usage)\n"
	ErrMsgEmptyInputFilename = "empty input filename"
	ErrMsgEmptyOutputFilename = "empty output filename"
	ErrMsgInputFileNotFound = "input file not found"
	ErrMsgOutputFileNotFound = "output file not found"
	ErrMsgOutputFileExists = "output file exists (use --force to overwrite)"
	)
