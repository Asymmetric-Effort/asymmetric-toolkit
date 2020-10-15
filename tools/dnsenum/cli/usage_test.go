package cli

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"crypto/sha256"
	"fmt"
	"testing"
)


func TestUsageConstantNotEmpty(t *testing.T) {
	errors.Assert(Usage != "", "expect non-empty string")
}

func TestUsageConstantHash(t *testing.T){
	//Hint: If this test fails, verify the usage.go content is correct
	//      then capture the actual hash from the test failure to update
	//      the expectedHash string below.
	const expectedHash = "9315e7a34a3c8e50131101f7a8de7a760a422b4363576855bd1acdc33a61c516"

	var usageHash = fmt.Sprintf("%x", sha256.Sum256([]byte(Usage)))
	errors.Assert(usageHash == expectedHash,
		fmt.Sprintf("Usage string does not match hash.\n"+
			"expected: %s\n" +
			"  actual: %s\n",
			expectedHash,
			usageHash))
}