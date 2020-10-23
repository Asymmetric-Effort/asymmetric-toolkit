package misc

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"fmt"
	"strconv"
	"testing"
)

func TestMiscIsPort(t *testing.T){
	for i:=0;i<65536;i++{
		errors.Assert(IsPort(strconv.Itoa(i)),fmt.Sprintf("Expected port (%d) to be valid",i))
	}
	errors.Assert(!IsPort("-1"),fmt.Sprintf("Expected port >= 0"))
	errors.Assert(!IsPort("65536"),fmt.Sprintf("Expected port <=65535"))
}