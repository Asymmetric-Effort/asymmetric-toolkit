package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestSpecification_ShowUsage(t *testing.T) {
	var o Specification

	fmt.Println("Starting TestSpecification_ShowVersion()")

	o.Version = "1.1.1"
	o.ProgramName = "program"
	o.Copyright = "copyright"
	o.Argument = make(map[string]ArgumentDescriptor)
	o.Argument["myFlag1"] = ArgumentDescriptor{
		1000,
		String,
		"myDefault1",
		"MyHelpString1",
		nil,
		ExpectNone,
	}
	o.Argument["myFlag2"] = ArgumentDescriptor{
		1000,
		String,
		"myDefault2",
		"MyHelpString2",
		nil,
		ExpectNone,
	}
	output := errors.CaptureStdOut(func() {
		inp := "--help"
		_, err := o.ShowUsage(&inp)
		if err != nil {
			panic(err)
		}
	})

	expected := fmt.Sprintf(bannerFmt, o.ProgramName, o.Version, o.Copyright) +
		fmt.Sprintf(lineFmt, "myFlag1", "MyHelpString1", "myDefault1") +
		fmt.Sprintf(lineFmt, "myFlag2", "MyHelpString2", "myDefault2")

	fmt.Printf("---START---\nERROR: Usage mismatch\n\n"+
		"Expected:\n%s\n"+
		"Actual:\n%s\n---DONE---\n",
		expected, output)

	fmt.Printf("\n"+
		"ExpectedLen: %d\n"+
		"ActualLen:   %d\n",
		len(expected), len(output))

	errors.Assert(output == expected, "Usage mismatch")
}
