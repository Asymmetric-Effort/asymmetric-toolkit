package entropy

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"github.com/google/uuid"
	"testing"
)

func TestHighEntropy(t *testing.T) {
	errors.Assert(!HighEntropy("1111111111111111111111"), "Expected false")
	errors.Assert(!HighEntropy("0"), "Expected false")
	errors.Assert(!HighEntropy("false"), "Expected false")
	errors.Assert(!HighEntropy("668108162888"), "Expected false")
	errors.Assert(HighEntropy(func()string{u,_:=uuid.NewUUID();return u.String()}()), "Expected true")
	errors.Assert(HighEntropy("ZmYwOTZmNmQyNWFjMWY4ZGY4MDBjNjQ3N2IwOGMxMDY4NTE1ODFjMjhlZmRjZGNmZmE2ZTM2MTQ4NjA2YTFkNDM2MDljZjc1MDFhODgxOTI0NGZmMmNmNmE1NWEyNDEzNmJjMWQxZmVkMmUwZmQ4ZDc5ODdiMjhiNzU4ZWUzYWYK"), "Expected true")
}
