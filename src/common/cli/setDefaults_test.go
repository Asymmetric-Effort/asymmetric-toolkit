package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestCommandLine_SetDefaults(t *testing.T) {
	func() { //test AddUsage only
		var o CommandLine
		var spec Specification
		fmt.Println("Test 1: --help ...starting")
		spec.AddUsage()
		err := o.SetDefaults(&spec) //Set defaults in the specification.
		if err != nil {
			panic(err)
		}
		fmt.Println("Test 1: --help ...passes")
	}()


	func() { //test AddDebug only
		var o CommandLine
		var spec Specification
		fmt.Println("Test 2 --debug ...starting")
		spec.AddDebug()
		err := o.SetDefaults(&spec) //Set defaults in the specification.
		if err != nil {
			panic(err)
		}
		errors.Assert(o.Arguments[flagDebug].Value=="false","Expect debug false.  Got "+o.Arguments[flagDebug].Value)
		errors.Assert(!o.Arguments[flagDebug].Boolean(),"Expect debug false.  Got "+o.Arguments[flagDebug].Value)
		fmt.Println("Test 2 --debug ...passes")
	}()


	func() { //test AddForce only
		var o CommandLine
		var spec Specification
		fmt.Println("Test 3")
		spec.AddForce()
		err := o.SetDefaults(&spec) //Set defaults in the specification.
		if err != nil {
			panic(err)
		}
		errors.Assert(o.Arguments[flagForce].Value=="false","Expect force false.  Got "+o.Arguments[flagForce].Value)
		errors.Assert(!o.Arguments[flagForce].Boolean(),"Expect force false.  Got "+o.Arguments[flagForce].Value)
		fmt.Println("Test 3 --force ...passes")
	}()


	func() { //test AddVersion only
		var o CommandLine
		var spec Specification
		fmt.Println("Test 4")
		spec.AddVersion()
		err := o.SetDefaults(&spec) //Set defaults in the specification.
		if err != nil {
			panic(err)
		}
		fmt.Println("Test 4 --version ...passes")
	}()

	func() { //Test all at once.
		var o CommandLine
		var spec Specification
		fmt.Println("Test 5")
		spec.AddUsage()
		spec.AddDebug()
		spec.AddForce()
		spec.AddVersion()

		err := o.SetDefaults(&spec) //Set defaults in the specification.
		if err != nil {
			panic(err)
		}
		errors.Assert(o.Arguments[flagDebug].Value=="false","Expect debug false.  Got "+o.Arguments[flagDebug].Value)
		errors.Assert(!o.Arguments[flagDebug].Boolean(),"Expect debug false.  Got "+o.Arguments[flagDebug].Value)
		errors.Assert(o.Arguments[flagForce].Value=="false","Expect force false.  Got "+o.Arguments[flagForce].Value)
		errors.Assert(!o.Arguments[flagForce].Boolean(),"Expect force false.  Got "+o.Arguments[flagForce].Value)
		fmt.Println("Test 5 ALL ...passes")
	}()
	/*

	 */
}
