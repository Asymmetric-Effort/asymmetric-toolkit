package entropy

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"github.com/google/uuid"
	"testing"
)

func TestGetShannons(t *testing.T){
	errors.Assert(GetShannons("0")==0,"Expected 0")
	errors.Assert(GetShannons("1111111111111111111111")==0,"Expected 0")
	errors.Assert(GetShannons(func()string{u,_:=uuid.NewUUID();return u.String()}())==144,"Expected 144")
}