package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
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
	const expectedHash = "868d2e41e44215aa64825275b67a697242ee3f02023f16a5ebd67dd6009e9184"

	var usageHash = fmt.Sprintf("%x", sha256.Sum256([]byte(Usage)))
	errors.Assert(usageHash == expectedHash,
		fmt.Sprintf("Usage string does not match hash.\n"+
			"expected: %s\n" +
			"  actual: %s\n",
			expectedHash,
			usageHash))
}