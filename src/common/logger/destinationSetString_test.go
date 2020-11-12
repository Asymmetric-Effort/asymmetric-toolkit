package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"testing"
)

func TestDestination_SetString(t *testing.T) {
	var o Destination

	o.SetString("stdout")
	errors.Assert(o.Get()== Stdout, "expected stdout")

	o.SetString("stdout")
	errors.Assert(o.Get()== Stdout, "expected stdout")

	o.SetString("stdout")
	errors.Assert(o.Get()== Stdout, "expected stdout")

	o.SetString("file")
	errors.Assert(o.Get()== File, "expected File")

	o.SetString("file")
	errors.Assert(o.Get()== File, "expected File")

	o.SetString("file")
	errors.Assert(o.Get()== File, "expected File")

	o.SetString("syslog")
	errors.Assert(o.Get()== Syslog, "expected syslog")

	o.SetString("syslog")
	errors.Assert(o.Get()== Syslog, "expected syslog")

	o.SetString("syslog")
	errors.Assert(o.Get()== Syslog, "expected syslog")
}
