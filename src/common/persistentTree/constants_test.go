package persistentTree

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestConstants(t *testing.T) {
	errors.Assert(fileHeaderVersion == 0, "Expected fileHeaderVersion == 0 --update test if intentional")
	errors.Assert(fileHeaderFlagCompressionEnabled == 1, "Flag value for enabling compression is 1")
	errors.Assert(fileHeaderFlagEncryptionEnabled == 2, "Flag value for enabling encryption is 2")
	errors.Assert(filePermissions == 0600, "File permissions should be 0600 allowing only owner read-write")
}
