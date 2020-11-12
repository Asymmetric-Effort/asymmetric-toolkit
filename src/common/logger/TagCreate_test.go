package logger

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"fmt"
	"testing"
)

func TestLogger_TagCreate(t *testing.T) {
	var o Logger
	o.Level = Debug
	o.Destination = Stdout
	o.Writer = func(p *string) {}
	t0 := o.nextTag
	t1 := o.TagCreate("tag")
	fmt.Println("t0:", t0, "t1", t1)
	errors.Assert(t0 < t1, "tag is not the next sequence.")
}
