package cli

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestSpecification(t *testing.T){
	var o Specification
	errors.Assert(o.Author == "", "expect empty string")
	errors.Assert(o.AuthorEmail == "", "expect empty string")
	errors.Assert(o.ProgramName == "", "expect empty string")
	errors.Assert(o.Copyright == "", "expect empty string")
	errors.Assert(o.Version == "", "expect empty string")
	errors.Assert(o.Argument == nil, "expect nil")
}
