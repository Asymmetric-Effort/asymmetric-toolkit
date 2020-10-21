package misc_test

import (
	"asymmetric-effort/asymmetric-toolkit/src/common/errors"
	"asymmetric-effort/asymmetric-toolkit/src/common/misc"
	"fmt"
	"testing"
)

func TestCharCountFrequency(t *testing.T){
	in:="AABBAACCDDD1111111112211"
	out:= misc.CountCharacterFrequency(&in)
	errors.Assert(int((*out)['A'])==4, fmt.Sprintf("Expect 2 A.  Encounter %d",int((*out)['A'])))
	errors.Assert(int((*out)['B'])==2, fmt.Sprintf("Expect 2 B.  Encounter %d",int((*out)['B'])))
	errors.Assert(int((*out)['C'])==2, fmt.Sprintf("Expect 2 C.  Encounter %d",int((*out)['C'])))
	errors.Assert(int((*out)['D'])==3, fmt.Sprintf("Expect 2 D.  Encounter %d",int((*out)['D'])))
	errors.Assert(int((*out)['1'])==11, fmt.Sprintf("Expect 2 D.  Encounter %d",int((*out)['D'])))
	errors.Assert(int((*out)['2'])==2, fmt.Sprintf("Expect 2 D.  Encounter %d",int((*out)['D'])))
}