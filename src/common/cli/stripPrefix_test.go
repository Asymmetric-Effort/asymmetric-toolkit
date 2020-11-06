package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestStripPrefix(t *testing.T) {

	func() {
		s := "--testarg"
		i, f := stripPrefix(&s)
		errors.Assert(i, "isFlag expected1")
		errors.Assert(f == "testarg", "long arg mismatch.")
	}()
	func() {
		s := "--ta"
		i, f := stripPrefix(&s)
		errors.Assert(i, "isFlag expected1")
		errors.Assert(f == "ta", "long arg mismatch.")
	}()
	func() {
		s := "-h"
		i, f := stripPrefix(&s)
		errors.Assert(i, "isFlag expected2")
		errors.Assert(f == "h", "short arg mismatch")
	}()
	func() {
		s:="badArg"
		i,f :=stripPrefix(&s)
		errors.Assert(!i, "isFlag not expected for badArg")
		errors.Assert(f==s, "badArg mismatch.1")
	}()

	func() {
		s := "--testA!rg"
		i, f := stripPrefix(&s)
		errors.Assert(!i, "isFlag not expected4")
		errors.Assert(f == "testA!rg", "badArg Mismatch2.")
	}()

	func(){
		s:="a"
		arg:=""
		for i:=0; len(arg)<24;i++ {
			s = fmt.Sprintf("%s%d",s,i)
			arg=fmt.Sprintf("--%s",s)
			i, f := stripPrefix(&arg)
			errors.Assert(i, fmt.Sprintf("isFlag expected for length test (long):%s",arg))
			errors.Assert(f == arg[2:], fmt.Sprintf("mismatch on length test (long):%s",arg))
		}
	}()

	func(){
		s:="abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
		i, f := stripPrefix(&s)
		errors.Assert(!i, "isFlag not expected (too long)")
		errors.Assert(f == "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "badArg Mismatch (toolong).")

	}()
}
