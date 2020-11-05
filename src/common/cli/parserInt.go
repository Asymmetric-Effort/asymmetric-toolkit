package cli

import (
	"fmt"
	"math"
	"strconv"
)

/*
	ParserInt() is a parser function factory which returns a parser function to the caller.
	This parser function will evaluate a given commandline argument (arg *string) and perform
	any validation before returning an appropriate error and Argument object pointer.
*/

func ParserInt(p ...int) (parser func(arg *string) (err error, val *Argument)) {
	/*
		Create the parser function and return it at the end.

		The parser will validate the floating point against any parameters (p) which are formatted as--
			p[0] ...minimum
			p[1] ...maximum
		to test inclusive bounds.
	*/
	const minimum = 0 // Index for p ...int (minimum bounds)
	const maximum = 1 // Index for p ...int (maximum bounds)

	getLimit := func(i int, defaultValue int) int {
		//
		// Given an array of parameter (p) values,
		//		where p[0] represents the minimum bound
		// 			and p[1] represents the maximum bound,
		//				and where we will use the given defaultValue where no p-value is provided,
		//		return the Limit to be applied when evaluating a commandline argument.
		//
		if len(p) >= (i + 1) {
			// A p-value is provided.  Return it.
			return p[i]
		} else {
			// No p-value is provided.  Return defaultValue.
			return defaultValue
		}
	}

	func() {
		//
		// Our parser function factory will now validate the bounds check parameters
		// using the smallest and largest data elements for the given type on the current platform
		// This ensures that the minimum bound is always less than the maximum bound BEFORE we use it.
		//
		min := getLimit(minimum, math.MinInt32)
		max := getLimit(maximum, math.MaxInt32)
		if min > max {
			panic("Specification contains float64 bounds where min > max.  Min must be less than max bound.")
		}
	}() //Note: we have executed the bounds check and found it works.

	parser = func(arg *string) (err error, val *Argument) {
		//
		// When our parser function runs, we validate whether the given commandline argument
		// matches parameter passed to the parser function in the Specification.
		//
		var value int // Our converted value.
		//
		//	Start off by converting our string commandline argument to its numeric equivalent.
		//
		value, err = strconv.Atoi(*arg)
		if err != nil {
			//
			// If our commandline argument cannot be parsed into a numeric value, return the
			// error and a nil Argument.
			//
			return err, nil
		}
		//
		// Get our actual upper and lower limits for the bounds check.  If none are provided, use the
		// numeric value of our commandline argument.
		//
		min := getLimit(minimum, value)
		max := getLimit(maximum, value)
		//
		// Perform out bounds check to ensure min <= value <= max.
		//
		if (value >= min) && (value <= max) {
			//
			// Bounds check okay, return the processed argument as a string to be extracted later.
			//
			return nil, &Argument{
				Float,
				*arg,
			}
		} else {
			//
			// Bound check failed.  Return an error and nil Argument object.
			//
			return fmt.Errorf("bounds check fail.  " +
				"Value (%d) falls outside the bound (%d...%d)", value, min, max), nil
		}
	}
	//return the parser function to the Specification.
	return parser
}
