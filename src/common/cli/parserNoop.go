package cli

/*
	The ParserNoop function is intended for use by programs which desire to do nothing when a flag is passed.  This is
	a placeholder useful for flags that have no value (e.g. --help, -h, -v or --version).  It will return a nil
	error and Argument and literally do nothing.  This comment does more.
*/

func ParserNoop() (parser func(arg *string) (err error, val *Argument)) {
	return func(arg *string)(err error, val *Argument) {
		return nil, nil
	}
}
